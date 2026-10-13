package repo

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	apperror "github.com/kodacampmain/koda-b9-gin/internal/error"
	"github.com/kodacampmain/koda-b9-gin/internal/model"
)

type PersonRepo struct {
}

func NewPersonRepo() *PersonRepo {
	return &PersonRepo{}
}

func (p *PersonRepo) GetAllPerson(ctx context.Context, db DBTX, limit, offset int) ([]model.Person, error) {
	// Jalankan Query
	sql := "SELECT name, age FROM persons LIMIT $1 OFFSET $2"
	args := []any{limit, offset}
	rows, err := db.Query(ctx, sql, args...)
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

func (P *PersonRepo) GetAllPersonMetadata(ctx context.Context, db DBTX) (int, error) {
	sql := "SELECT COUNT(id) FROM persons"
	var count int
	if err := db.QueryRow(ctx, sql).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (p *PersonRepo) GetPersonById(ctx context.Context, db DBTX, id int) (model.Person, error) {
	// SQL injection
	// sql := fmt.Sprintf("SELECT id, name, age FROM persons WHERE id=%s", id)
	sql := "SELECT id, name, age FROM persons WHERE id=$1"
	// $ => Parameterized Query untuk PSQL
	// Parameterized Query di psql dibaca berdasarkan posisi
	args := []any{id}
	row := db.QueryRow(ctx, sql, args...)
	var person model.Person
	if err := row.Scan(&person.Id, &person.Name, &person.Age); err != nil {
		// Expected Error
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Person{}, apperror.ErrNoData
		}
		// Unexpected Error
		return model.Person{}, err
	}
	return person, nil
}

func (p *PersonRepo) CreateNewPerson(ctx context.Context, db DBTX, newPerson model.Person) error {
	sql := "INSERT INTO persons (name, age) VALUES ($1,$2)"
	// sql := "INSERT INTO persons (name, age) VALUES ($1,$2) RETURNING id, name, age"
	// krn returning maka ada row result, jadi gunakan query atau queryrow berdasarkan ekspektasi output
	args := []any{newPerson.Name, newPerson.Age}
	cmd, err := db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("no change")
	}
	return nil
}
