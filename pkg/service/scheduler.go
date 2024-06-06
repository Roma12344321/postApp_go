package service

import (
	"context"
	"fmt"
	"log"
	"postApp/pkg/repository"
	"time"
)

const (
	vipRole          = "ROLE_VIP"
	secondsForUpdate = 20
	factor           = 1.1
)

type Scheduler struct {
	repo *repository.Repository
}

func NewScheduler(repo *repository.Repository) *Scheduler {
	return &Scheduler{repo: repo}
}

func (s *Scheduler) UpdateBalanceForVipPerson(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Scheduler is shutting down")
			return
		case <-time.After(secondsForUpdate * time.Second):
			people, err := s.repo.PersonRepository.GetAllPersonWithRole(vipRole)
			if err != nil {
				log.Fatalln(err)
			}
			err = s.repo.BalanceRepository.UpdateBalanceForPeople(factor, &people)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
