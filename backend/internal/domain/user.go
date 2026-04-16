package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	IDToko         string         `gorm:"type:uuid;not null" json:"id_toko"`
	Nama           string         `gorm:"type:text;not null" json:"nama"`
	Email          string         `gorm:"type:text;unique;not null" json:"email"`
	Password       string         `gorm:"type:text;not null" json:"-"`
	Role           string         `gorm:"type:user_role;default:'kasir'" json:"role"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
