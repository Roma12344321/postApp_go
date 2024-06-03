package repository

import (
	"github.com/jmoiron/sqlx"
	"postApp"
)

type AuthRepository interface {
	CreatePerson(person *postApp.Person) (int, error)
}

type PersonRepository interface {
	GetPerson(username, password string) (*postApp.Person, error)
}

type Repository struct {
	AuthRepository
	PersonRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthRepositoryImpl(db), NewPersonRepositoryImpl(db)}
}
