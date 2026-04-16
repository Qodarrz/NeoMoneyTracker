package controller

import (
	"net/http"

	"github.com/Qodarrz/go-gin-air/internal/dto"
	"github.com/Qodarrz/go-gin-air/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	service *service.AdminService
}

func NewAdminController(service *service.AdminService) *AdminController {
	return &AdminController{service: service}
}

func (c *AdminController) getIDToko(ctx *gin.Context) string {
	id, _ := ctx.Get("id_toko")
	return id.(string)
}

// --- Toko ---

func (c *AdminController) UpdateToko(ctx *gin.Context) {
	var req dto.TokoUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateToko(c.getIDToko(ctx), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Shop updated successfully"})
}

// --- Staff ---

func (c *AdminController) ListStaff(ctx *gin.Context) {
	staff, err := c.service.ListStaff(c.getIDToko(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, staff)
}

func (c *AdminController) CreateKasir(ctx *gin.Context) {
	var req dto.CreateKasirRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateKasir(c.getIDToko(ctx), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Staff created successfully"})
}

func (c *AdminController) PromoteToAdmin(ctx *gin.Context) {
	var req dto.PromoteRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.PromoteToAdmin(c.getIDToko(ctx), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Staff promoted to admin successfully"})
}

// --- Products ---

func (c *AdminController) ListProducts(ctx *gin.Context) {
	products, err := c.service.ListProducts(c.getIDToko(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c *AdminController) CreateProduct(ctx *gin.Context) {
	var req dto.ProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateProduct(c.getIDToko(ctx), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func (c *AdminController) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	var req dto.ProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateProduct(c.getIDToko(ctx), productID, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (c *AdminController) DeleteProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	if err := c.service.DeleteProduct(c.getIDToko(ctx), productID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
