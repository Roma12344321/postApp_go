package service

import (
	"crypto/sha1"
	"fmt"
	"postApp"
	"postApp/pkg/repository"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

type AuthServiceImpl struct {
	repo *repository.Repository
}

func NewAuthServiceImpl(repo *repository.Repository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) CreatePerson(person *postApp.Person) (int, error) {
	passwordHash := generatePasswordHash(person.Password)
	person.Password = passwordHash
	return s.repo.AuthRepository.CreatePerson(person)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
