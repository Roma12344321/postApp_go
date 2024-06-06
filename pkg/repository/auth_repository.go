package repository

import (
	"github.com/jmoiron/sqlx"
	"postApp"
)

type AuthRepositoryImpl struct {
	db *sqlx.DB
}

func NewAuthRepositoryImpl(db *sqlx.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) CreatePerson(person *postApp.Person) (int, error) {
	query := "INSERT INTO person(username, password,role) VALUES ($1,$2,$3) RETURNING id"
	row := r.db.QueryRow(query, person.Username, person.Password, person.Role)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
