package service2

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// RepositoryImpl ...
type RepositoryImpl struct {
	db *sqlx.DB
}

// NewRepositoryImpl ...
func NewRepositoryImpl(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

// RegisterPerson ...
func (r *RepositoryImpl) RegisterPerson(ctx context.Context, p Person) (Person, error) {
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
func (r *RepositoryImpl) GetPersonByID(ctx context.Context, id int64) (Person, error) {
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
