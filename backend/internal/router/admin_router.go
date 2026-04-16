package router

import (
	"github.com/Qodarrz/go-gin-air/config"
	"github.com/Qodarrz/go-gin-air/internal/controller"
	"github.com/Qodarrz/go-gin-air/internal/middleware"
	"github.com/Qodarrz/go-gin-air/internal/repository"
	"github.com/Qodarrz/go-gin-air/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(r *gin.Engine) {
	db := config.DB
	tokoRepo := repository.NewTokoRepository(db)
	userRepo := repository.NewUserRepository(db)
	produkRepo := repository.NewProdukRepository(db)

	svc := service.NewAdminService(tokoRepo, userRepo, produkRepo)
	ctrl := controller.NewAdminController(svc)

	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RoleMiddleware("admin"))
	{
		// Toko
		admin.PUT("/toko", ctrl.UpdateToko)

		// Staff
		admin.GET("/staff", ctrl.ListStaff)
		admin.POST("/staff", ctrl.CreateKasir)
		admin.POST("/staff/promote", ctrl.PromoteToAdmin)

		// Products
		admin.GET("/products", ctrl.ListProducts)
		admin.POST("/products", ctrl.CreateProduct)
		admin.PUT("/products/:id", ctrl.UpdateProduct)
		admin.DELETE("/products/:id", ctrl.DeleteProduct)
	}
}
