package repository

import (
	"github.com/Qodarrz/go-gin-air/internal/domain"
	"gorm.io/gorm"
)

type ProdukRepository interface {
	Create(produk *domain.Produk) error
	Update(produk *domain.Produk) error
	Delete(id string, idToko string) error
	FindByID(id string, idToko string) (*domain.Produk, error)
	FindAll(idToko string) ([]domain.Produk, error)
}

type gormProdukRepository struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &gormProdukRepository{db: db}
}

func (r *gormProdukRepository) Create(produk *domain.Produk) error {
	return r.db.Create(produk).Error
}

func (r *gormProdukRepository) Update(produk *domain.Produk) error {
	return r.db.Save(produk).Error
}

func (r *gormProdukRepository) Delete(id string, idToko string) error {
	return r.db.Where("id = ? AND id_toko = ?", id, idToko).Delete(&domain.Produk{}).Error
}

func (r *gormProdukRepository) FindByID(id string, idToko string) (*domain.Produk, error) {
	var produk domain.Produk
	if err := r.db.Where("id = ? AND id_toko = ?", id, idToko).First(&produk).Error; err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *gormProdukRepository) FindAll(idToko string) ([]domain.Produk, error) {
	var produks []domain.Produk
	if err := r.db.Where("id_toko = ?", idToko).Find(&produks).Error; err != nil {
		return nil, err
	}
	return produks, nil
}
