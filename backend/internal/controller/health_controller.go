package controller

import (
	"net/http"

	"github.com/Qodarrz/go-gin-air/internal/service"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	service *service.HealthService
}

func NewHealthController(service *service.HealthService) *HealthController {
	return &HealthController{
		service: service,
	}
}

func (c *HealthController) Check(ctx *gin.Context) {
	status := c.service.Check()

	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
