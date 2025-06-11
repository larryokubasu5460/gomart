package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/larryokubasu5460/product-service/model"
	"github.com/larryokubasu5460/product-service/service"
)

type ProductHandler struct {
	Service service.ProductService
}

func (h *ProductHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/products", h.Create)
	r.GET("/products", h.GetAll)
	r.GET("/products:id", h.GetByID)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	if err := h.Service.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetAll(c *gin.Context){
	products, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid ID"})
		return
	}
	product, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}