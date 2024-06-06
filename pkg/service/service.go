package service

import (
	"context"
	"postApp"
	"postApp/pkg/repository"
)

type AuthService interface {
	Registration(person *postApp.Person) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type SchedulerService interface {
	UpdateBalanceForVipPerson(ctx context.Context)
}

type Service struct {
	AuthService
	SchedulerService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{NewAuthServiceImpl(repository), NewScheduler(repository)}
}
