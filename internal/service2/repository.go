package service2

import "context"

// Repository ...
type Repository interface {
	RegisterPerson(context.Context, Person) (Person, error)
	GetPersonByID(context.Context, int64) (Person, error)
}
