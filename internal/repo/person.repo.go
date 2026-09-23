package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b9-gin/internal/model"
)

type PersonRepo struct {
	db *pgxpool.Pool
}

func NewPersonRepo(db *pgxpool.Pool) *PersonRepo {
	return &PersonRepo{
		db: db,
	}
}

func (p *PersonRepo) GetAllPerson(ctx context.Context) ([]model.Person, error) {
	// Jalankan Query
	sql := "SELECT name, age FROM persons"
	rows, err := p.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	// Olah hasil query
	var persons []model.Person
	for rows.Next() {
		var person model.Person
		if err := rows.Scan(&person.Name, &person.Age); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	// return data jika berhasil
	return persons, nil
}
