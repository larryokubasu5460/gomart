package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/larryokubasu5460/product-service/config"
	"github.com/larryokubasu5460/product-service/handler"
	"github.com/larryokubasu5460/product-service/repository"
	"github.com/larryokubasu5460/product-service/service"
)

func main() {
	godotenv.Load()

	db := config.InitDB()
	r := gin.Default()

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.ProductHandler{Service: productService}
	productHandler.RegisterRoutes(r)

	r.Run(":8081")
}