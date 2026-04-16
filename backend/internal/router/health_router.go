package router

import (
	"github.com/Qodarrz/go-gin-air/internal/controller"
	"github.com/Qodarrz/go-gin-air/internal/repository"
	"github.com/Qodarrz/go-gin-air/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterHealthRouter(r *gin.Engine) {
	// dependency wiring
	repo := repository.NewHealthRepository()
	svc := service.NewHealthService(repo)
	ctrl := controller.NewHealthController(svc)

	r.GET("/health/check", ctrl.Check)
}
