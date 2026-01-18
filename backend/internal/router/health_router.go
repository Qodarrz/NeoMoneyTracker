package router

import (
	"net/http"

	"github.com/Qodarrz/go-gin-air/internal/controller"
	"github.com/Qodarrz/go-gin-air/internal/repository"
	"github.com/Qodarrz/go-gin-air/internal/service"
)

func RegisterHealthRouter(mux *http.ServeMux) {
	// dependency wiring
	repo := repository.NewHealthRepository()
	svc := service.NewHealthService(repo)
	ctrl := controller.NewHealthController(svc)

	mux.HandleFunc("/health/check", ctrl.Check)
}
