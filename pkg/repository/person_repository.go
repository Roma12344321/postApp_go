package repository

import (
	"github.com/jmoiron/sqlx"
	"postApp"
)

type PersonRepositoryImpl struct {
	db *sqlx.DB
}

func NewPersonRepositoryImpl(db *sqlx.DB) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{db: db}
}

func (r *PersonRepositoryImpl) GetPerson(username, password string) (*postApp.Person, error) {
	var person postApp.Person
	query := "SELECT * FROM person WHERE username=$1 AND password=$2"
	if err := r.db.Get(&person, query, username, password); err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *PersonRepositoryImpl) GetAllPersonWithRole(role string) ([]postApp.Person, error) {
	var people []postApp.Person
	query := "SELECT * FROM person WHERE role=$1"
	if err := r.db.Select(&people, query, role); err != nil {
		return nil, err
	}
	return people, nil
}
