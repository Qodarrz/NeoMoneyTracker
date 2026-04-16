package dto

type RegisterRequest struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
	IDToko   string `json:"id_toko" form:"id_toko" binding:"required"`
	Role     string `json:"role" form:"role"` // optional, defaults to kasir
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type AuthResponse struct {
	Token    string `json:"token"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
