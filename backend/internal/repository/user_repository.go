package repository

import (
	"github.com/Qodarrz/go-gin-air/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *domain.User) error
	Update(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id string, idToko string) (*domain.User, error)
	FindAllByToko(idToko string) ([]domain.User, error)
	ExistsByEmail(email string) bool
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{
		db: db,
	}
}

func (r *gormUserRepository) Save(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *gormUserRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *gormUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) FindByID(id string, idToko string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ? AND id_toko = ?", id, idToko).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) FindAllByToko(idToko string) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Where("id_toko = ?", idToko).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *gormUserRepository) ExistsByEmail(email string) bool {
	var count int64
	r.db.Model(&domain.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
