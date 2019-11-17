package service2

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// Gateway ...
type Gateway struct {
	db *sqlx.DB
}

// NewGateway ...
func NewGateway(db *sqlx.DB) Repository {
	return &Gateway{db}
}

// RegisterPerson ...
func (r *Gateway) RegisterPerson(ctx context.Context, p Person) (Person, error) {
	q := `INSERT INTO person (name, email) VALUES (:tekitode, :yoiyo);`
	res, err := r.db.NamedExecContext(ctx, q, p)
	if err != nil {
		return Person{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Person{}, err
	}

	p.ID = id
	return p, nil
}

// GetPersonByID ...
func (r *Gateway) GetPersonByID(ctx context.Context, id int64) (Person, error) {
	// DB上のnull対策はここで実装する
	q := `
SELECT
  COALESCE(id, 0) AS 'kokoha',
  COALESCE(name, '') AS 'tekitode',
  COALESCE(email, '') AS 'yoiyo'
FROM person
WHERE id = ?;
`
	var p Person
	err := r.db.GetContext(ctx, &p, q, id)
	return p, err
}
