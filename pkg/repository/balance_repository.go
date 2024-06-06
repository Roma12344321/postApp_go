package repository

import (
	"github.com/jmoiron/sqlx"
	"postApp"
)

type BalanceRepositoryImpl struct {
	db *sqlx.DB
}

func NewBalanceRepositoryImpl(db *sqlx.DB) *BalanceRepositoryImpl {
	return &BalanceRepositoryImpl{db: db}
}

func (r *BalanceRepositoryImpl) CreateBalance(balance *postApp.Balance) {
	query := "INSERT INTO balance VALUES ($1,$2)"
	r.db.QueryRow(query, balance.PersonId, balance.Sum)
}

func (r *BalanceRepositoryImpl) UpdateBalanceForPeople(factor float64, people *[]postApp.Person) error {
	for _, p := range *people {
		query := "UPDATE balance SET sum=sum*$1 WHERE person_id=$2"
		_, err := r.db.Exec(query, factor, p.Id)
		if err != nil {
			return err
		}
	}
	return nil
}
