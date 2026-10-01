package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/internal/repo"
	"github.com/redis/go-redis/v9"
)

type PersonService struct {
	pr  *repo.PersonRepo
	rdb *redis.Client
}

func NewPersonService(pr *repo.PersonRepo, rdb *redis.Client) *PersonService {
	return &PersonService{
		pr:  pr,
		rdb: rdb,
	}
}

func (p *PersonService) GetAllPerson(ctx context.Context) ([]dto.Person, error) {
	// cek ke redis
	key := "fakhridho:persons"

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
			return persons, nil
		}
	}
	// cache miss, ambil dari db
	result, err := p.pr.GetAllPerson(ctx)
	data := make([]dto.Person, 0, len(result))
	for _, v := range result {
		data = append(data, dto.Person{
			Name: v.Name,
			Age:  v.Age,
		})
	}

	// lalu simpan ke redis
	if str, err := json.Marshal(data); err != nil {
		log.Println("Stringify error\nReason: ", err.Error())
	} else {
		if err := p.rdb.Set(ctx, key, string(str), 0).Err(); err != nil {
			log.Println("Redis set error\nReason: ", err.Error())
		}
	}
	// kembalikan data ke handler/controller
	return data, err
}
