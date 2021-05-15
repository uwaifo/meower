package db

import (
	"context"
	"database/sql"

	"github.com/uwaifo/meower/schema"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()

}

func (r *PostgresRepository) InsertMeow(ctx context.Context, meow schema.Meow) error {
	_, err := r.db.Exec("INSERT INTO meows(id, body, created_at) VALUES($1, $2, $3)", meow.ID, meow.Body, meow.CreatedAt)

	return err
}

func (r *PostgresRepository) ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	rows, err := r.db.Query("SELECT * FROM meows ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//
	listMeows := []schema.Meow{}
	for rows.Next() {
		meow := schema.Meow{}
		if err := rows.Scan(&meow.ID, &meow.Body, &meow.CreatedAt); err != nil {
			listMeows = append(listMeows, meow)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listMeows, nil
}

//TODO Meows are being ordered by their primary key, because keys will be k-sortable by time.
//TODO This is to avoid introducing an additional index
