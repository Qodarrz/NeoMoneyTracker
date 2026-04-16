package router

import (
	"net/http"

	"github.com/Qodarrz/go-gin-air/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Global Middlewares (Logger, Recovery, etc.)
	// r.Use(gin.Logger())

	// Public Routes
	RegisterHealthRouter(r)
	RegisterAuthRouter(r)
	RegisterAdminRouter(r)

	// Protected Routes (Example)
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Admin Only Route Example
		admin := api.Group("/admin")
		admin.Use(middleware.RoleMiddleware("admin"))
		{
			admin.GET("/dashboard", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Admin Dashboard"})
			})
		}

		// Any Role Route Example
		api.GET("/profile", func(ctx *gin.Context) {
			userID, _ := ctx.Get("user_id")
			role, _ := ctx.Get("role")
			ctx.JSON(http.StatusOK, gin.H{
				"user_id": userID,
				"role":    role,
				"message": "This is your profile",
			})
		})
	}

	return r
}
