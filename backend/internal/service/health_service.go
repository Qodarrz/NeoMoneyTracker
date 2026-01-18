package service

import "github.com/Qodarrz/go-gin-air/internal/repository"

type HealthService struct {
	repo *repository.HealthRepository
}

func NewHealthService(repo *repository.HealthRepository) *HealthService {
	return &HealthService{
		repo: repo,
	}
}

func (s *HealthService) Check() string {
	health := s.repo.Check()
	return health.Status
}
