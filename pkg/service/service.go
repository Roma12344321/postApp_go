package service

import (
	"postApp"
	"postApp/pkg/repository"
)

type AuthService interface {
	CreatePerson(person *postApp.Person) (int,error)
}

type Service struct {
	AuthService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{NewAuthServiceImpl(repository)}
}
