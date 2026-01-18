package controller

import (
	"net/http"

	"github.com/Qodarrz/go-gin-air/internal/helper"
	"github.com/Qodarrz/go-gin-air/internal/service"
)

type HealthController struct {
	service *service.HealthService
}

func NewHealthController(service *service.HealthService) *HealthController {
	return &HealthController{
		service: service,
	}
}

func (c *HealthController) Check(w http.ResponseWriter, r *http.Request) {
	status := c.service.Check()

	helper.JSON(w, http.StatusOK, map[string]string{
		"status": status,
	})
}
