package repository

import (
	"github.com/jmoiron/sqlx"
	"postApp"
)

type AuthRepository interface {
	CreatePerson(person *postApp.Person) (int,error)
}

type Repository struct {
	AuthRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthRepositoryImpl(db)}
}