package domain

import (
	"time"

	"gorm.io/gorm"
)

type Toko struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	NamaToko  string         `gorm:"type:text;not null" json:"nama_toko"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
