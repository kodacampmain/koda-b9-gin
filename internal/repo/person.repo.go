package repo

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	apperror "github.com/kodacampmain/koda-b9-gin/internal/error"
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

func (p *PersonRepo) GetPersonById(ctx context.Context, id int) (model.Person, error) {
	// SQL injection
	// sql := fmt.Sprintf("SELECT id, name, age FROM persons WHERE id=%s", id)
	sql := "SELECT id, name, age FROM persons WHERE id=$1"
	// $ => Parameterized Query untuk PSQL
	// Parameterized Query di psql dibaca berdasarkan posisi
	args := []any{id}
	row := p.db.QueryRow(ctx, sql, args...)
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

func (p *PersonRepo) CreateNewPerson(ctx context.Context, newPerson model.Person) error {
	sql := "INSERT INTO persons (name, age) VALUES ($1,$2)"
	// sql := "INSERT INTO persons (name, age) VALUES ($1,$2) RETURNING id, name, age"
	// krn returning maka ada row result, jadi gunakan query atau queryrow berdasarkan ekspektasi output
	args := []any{newPerson.Name, newPerson.Age}
	cmd, err := p.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("no change")
	}
	return nil
}
