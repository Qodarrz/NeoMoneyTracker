package repository

import "github.com/Qodarrz/go-gin-air/internal/domain"

type HealthRepository struct{}

func NewHealthRepository() *HealthRepository {
	return &HealthRepository{}
}

func (r *HealthRepository) Check() domain.Health {
	return domain.Health{
		Status: "ok",
	}
}
