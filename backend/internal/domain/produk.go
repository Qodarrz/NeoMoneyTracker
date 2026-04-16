package domain

import (
	"time"

	"gorm.io/gorm"
)

type Produk struct {
	ID         string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	IDToko     string         `gorm:"type:uuid;not null" json:"id_toko"`
	Barcode    string         `gorm:"type:text;not null" json:"barcode"`
	NamaProduk string         `gorm:"type:text;not null" json:"nama_produk"`
	Harga      float64        `gorm:"type:decimal(15,2);not null;default:0" json:"harga"`
	Stok       int            `gorm:"type:int;not null;default:0" json:"stok"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
