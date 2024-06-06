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
	GetAllPersonWithRole(role string) ([]postApp.Person, error)
}

type BalanceRepository interface {
	CreateBalance(balance *postApp.Balance)
	UpdateBalanceForPeople(factor float64, people *[]postApp.Person) error
}

type Repository struct {
	AuthRepository
	PersonRepository
	BalanceRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthRepositoryImpl(db),
		NewPersonRepositoryImpl(db),
		NewBalanceRepositoryImpl(db)}
}
