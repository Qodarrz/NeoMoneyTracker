package repository

import (
	"github.com/Qodarrz/go-gin-air/internal/domain"
	"gorm.io/gorm"
)

type TokoRepository interface {
	Create(toko *domain.Toko) error
	Update(toko *domain.Toko) error
	FindByID(id string) (*domain.Toko, error)
}

type gormTokoRepository struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &gormTokoRepository{db: db}
}

func (r *gormTokoRepository) Create(toko *domain.Toko) error {
	return r.db.Create(toko).Error
}

func (r *gormTokoRepository) Update(toko *domain.Toko) error {
	return r.db.Save(toko).Error
}

func (r *gormTokoRepository) FindByID(id string) (*domain.Toko, error) {
	var toko domain.Toko
	if err := r.db.First(&toko, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &toko, nil
}
