package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/url"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/internal/repo"
	"github.com/redis/go-redis/v9"
)

type PersonService struct {
	pr  *repo.PersonRepo
	rdb *redis.Client
	db  *pgxpool.Pool
}

func NewPersonService(pr *repo.PersonRepo, rdb *redis.Client, db *pgxpool.Pool) *PersonService {
	return &PersonService{
		pr:  pr,
		rdb: rdb,
		db:  db,
	}
}

func (p *PersonService) GetAllPerson(ctx context.Context, endpoint string, qp dto.PersonQuery) ([]dto.Person, dto.PaginationMeta, error) {
	// cek ke redis
	key := fmt.Sprintf("fakhridho:persons:%d", qp.Page)
	totalDataKey := "fakhridho:persons:total"

	// key := "fakhridho:persons"
	// if qp.Page == 1 {
	if str, err := p.rdb.Get(ctx, key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			log.Println("key does not exist")
		} else {
			log.Println(err.Error())
		}
	} else {
		// cache hit, return output
		var persons []dto.Person
		if err := json.Unmarshal([]byte(str), &persons); err != nil {
			log.Println("Parse error\nReason: ", err.Error())
		} else {
			if totalData, err := p.rdb.Get(ctx, totalDataKey).Result(); err != nil {
				if errors.Is(err, redis.Nil) {
					log.Println("key does not exist")
				} else {
					log.Println(err.Error())
				}
			} else {
				td, err := strconv.Atoi(totalData)
				if err != nil {
					log.Println(err.Error())
				} else {
					totalPage := math.Ceil(float64(td) / float64(qp.Limit))

					var next string
					if qp.Page != int(totalPage) {
						query := url.Values{}
						query.Set("limit", strconv.Itoa(qp.Limit))
						query.Set("page", strconv.Itoa(qp.Page+1))
						next = fmt.Sprintf("%s?%s", endpoint, query.Encode())
					}

					var prev string
					if qp.Page != 1 {
						query := url.Values{}
						query.Set("limit", strconv.Itoa(qp.Limit))
						query.Set("page", strconv.Itoa(qp.Page-1))
						prev = fmt.Sprintf("%s?%s", endpoint, query.Encode())
					}

					meta := dto.PaginationMeta{
						CurrentPage: qp.Page,
						TotalPage:   totalPage,
						Next:        next,
						Prev:        prev,
					}
					return persons, meta, nil
				}
			}

		}
	}
	// }

	// cache miss, ambil dari db
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return []dto.Person{}, dto.PaginationMeta{}, err
	}
	defer tx.Rollback(ctx)

	offset := (qp.Page - 1) * qp.Limit
	result, err := p.pr.GetAllPerson(ctx, tx, qp.Limit, offset)
	data := make([]dto.Person, 0, len(result))
	for _, v := range result {
		data = append(data, dto.Person{
			Name: v.Name,
			Age:  v.Age,
		})
	}

	totalData, err := p.pr.GetAllPersonMetadata(ctx, tx)
	if err != nil {
		return []dto.Person{}, dto.PaginationMeta{}, err
	}

	totalPage := math.Ceil(float64(totalData) / float64(qp.Limit))

	var next string
	if qp.Page < int(totalPage) {
		query := url.Values{}
		query.Set("limit", strconv.Itoa(qp.Limit))
		query.Set("page", strconv.Itoa(qp.Page+1))
		next = fmt.Sprintf("%s?%s", endpoint, query.Encode())
	}

	var prev string
	if qp.Page > 1 && qp.Page <= int(totalPage) {
		query := url.Values{}
		query.Set("limit", strconv.Itoa(qp.Limit))
		query.Set("page", strconv.Itoa(qp.Page-1))
		prev = fmt.Sprintf("%s?%s", endpoint, query.Encode())
	}

	meta := dto.PaginationMeta{
		CurrentPage: qp.Page,
		TotalPage:   totalPage,
		Next:        next,
		Prev:        prev,
	}

	if len(data) > 0 {
		// if qp.Page == 1 {
		// lalu simpan ke redis
		if str, err := json.Marshal(data); err != nil {
			log.Println("Stringify error\nReason: ", err.Error())
		} else {
			if err := p.rdb.Set(ctx, key, string(str), 0).Err(); err != nil {
				log.Println("Redis set error\nReason: ", err.Error())
			}
		}
		// }
		if err := p.rdb.Set(ctx, totalDataKey, totalData, 0).Err(); err != nil {
			log.Println("Redis set error\nReason: ", err.Error())
		}
	}

	// kembalikan data ke handler/controller
	return data, meta, err
}
