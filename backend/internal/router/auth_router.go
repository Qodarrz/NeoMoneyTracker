package router

import (
	"github.com/Qodarrz/go-gin-air/config"
	"github.com/Qodarrz/go-gin-air/internal/controller"
	"github.com/Qodarrz/go-gin-air/internal/repository"
	"github.com/Qodarrz/go-gin-air/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRouter(r *gin.Engine) {
	// Using the global GORM DB
	db := config.DB
	repo := repository.NewUserRepository(db)
	svc := service.NewAuthService(repo)
	ctrl := controller.NewAuthController(svc)

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", ctrl.Register)
		auth.POST("/login", ctrl.Login)
	}
}
