package main

import (
	"net/http"

	"github.com/Qodarrz/go-gin-air/internal/router"
)

func main() {
	r := router.NewRouter()
	http.ListenAndServe(":8080", r)
}
