package service

import "github.com/mohsanabbas/loggi/internal/repository"

type Service interface {
	PrintInService(string) (string, error)
}
type service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) PrintInService(m string) (string, error) {
	result, err := s.repository.Print(m)
	return result, err
}
