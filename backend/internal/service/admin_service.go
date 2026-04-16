package service

import (
	"errors"

	"github.com/Qodarrz/go-gin-air/internal/domain"
	"github.com/Qodarrz/go-gin-air/internal/dto"
	"github.com/Qodarrz/go-gin-air/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	tokoRepo   repository.TokoRepository
	userRepo   repository.UserRepository
	produkRepo repository.ProdukRepository
}

func NewAdminService(
	tokoRepo repository.TokoRepository,
	userRepo repository.UserRepository,
	produkRepo repository.ProdukRepository,
) *AdminService {
	return &AdminService{
		tokoRepo:   tokoRepo,
		userRepo:   userRepo,
		produkRepo: produkRepo,
	}
}

// --- Toko Management ---

func (s *AdminService) UpdateToko(idToko string, req dto.TokoUpdateRequest) error {
	toko, err := s.tokoRepo.FindByID(idToko)
	if err != nil {
		return err
	}
	toko.NamaToko = req.NamaToko
	return s.tokoRepo.Update(toko)
}

// --- User (Staff) Management ---

func (s *AdminService) ListStaff(idToko string) ([]dto.UserWithTokoResponse, error) {
	users, err := s.userRepo.FindAllByToko(idToko)
	if err != nil {
		return nil, err
	}

	var res []dto.UserWithTokoResponse
	for _, u := range users {
		res = append(res, dto.UserWithTokoResponse{
			ID:    u.ID,
			Nama:  u.Nama,
			Email: u.Email,
			Role:  u.Role,
		})
	}
	return res, nil
}

func (s *AdminService) CreateKasir(idToko string, req dto.CreateKasirRequest) error {
	if s.userRepo.ExistsByEmail(req.Email) {
		return errors.New("email already exists")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := &domain.User{
		IDToko:   idToko,
		Nama:     req.Nama,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "kasir",
	}

	return s.userRepo.Save(user)
}

func (s *AdminService) PromoteToAdmin(idToko string, req dto.PromoteRequest) error {
	user, err := s.userRepo.FindByID(req.UserID, idToko)
	if err != nil {
		return errors.New("user not found in this shop")
	}

	user.Role = "admin"
	return s.userRepo.Update(user)
}

// --- Product Management ---

func (s *AdminService) ListProducts(idToko string) ([]domain.Produk, error) {
	return s.produkRepo.FindAll(idToko)
}

func (s *AdminService) CreateProduct(idToko string, req dto.ProductRequest) error {
	produk := &domain.Produk{
		IDToko:     idToko,
		Barcode:    req.Barcode,
		NamaProduk: req.NamaProduk,
		Harga:      req.Harga,
		Stok:       req.Stok,
	}
	return s.produkRepo.Create(produk)
}

func (s *AdminService) UpdateProduct(idToko string, productID string, req dto.ProductRequest) error {
	produk, err := s.produkRepo.FindByID(productID, idToko)
	if err != nil {
		return errors.New("product not found")
	}

	produk.Barcode = req.Barcode
	produk.NamaProduk = req.NamaProduk
	produk.Harga = req.Harga
	produk.Stok = req.Stok

	return s.produkRepo.Update(produk)
}

func (s *AdminService) DeleteProduct(idToko string, productID string) error {
	return s.produkRepo.Delete(productID, idToko)
}
