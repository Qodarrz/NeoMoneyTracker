package dto

type TokoUpdateRequest struct {
	NamaToko string `json:"nama_toko" form:"nama_toko" binding:"required"`
}

type CreateKasirRequest struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type ProductRequest struct {
	Barcode    string  `json:"barcode" form:"barcode" binding:"required"`
	NamaProduk string  `json:"nama_produk" form:"nama_produk" binding:"required"`
	Harga      float64 `json:"harga" form:"harga" binding:"required,min=0"`
	Stok       int     `json:"stok" form:"stok" binding:"required,min=0"`
}

type PromoteRequest struct {
	UserID string `json:"user_id" form:"user_id" binding:"required"`
}

type UserWithTokoResponse struct {
	ID    string `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
