package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/larryokubasu5460/order-service/client"
	"github.com/larryokubasu5460/order-service/config"
	"github.com/larryokubasu5460/order-service/db"
	"github.com/larryokubasu5460/order-service/handler"
	"github.com/larryokubasu5460/order-service/model"
	"github.com/larryokubasu5460/order-service/repository"
	"github.com/larryokubasu5460/order-service/service"
)

func main() {
	// // load .env
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found, using system env")
	// }

	// init config
	cfg := config.LoadConfig()

	// Connect to DB
	database := db.ConnectPostgres(cfg.DBUrl)

	_ = database.AutoMigrate(&model.Order{}, &model.OrderItem{})

	// Init repo, clients, services
	orderRepo := repository.NewOrderRepository(database)
	userClient := client.NewUserClient(cfg.UserServiceURL)
	productClient := client.NewProductClient(cfg.ProductServiceURL)

	orderService := service.NewOrderService(orderRepo, userClient, productClient)
	orderHandler := handler.NewOrderHandler(orderService)

	// Gin setup
	r := gin.Default()
	
	orderHandler.RegisterRoutes(r)

	log.Printf("Order service running on port %s", cfg.Port)
	if err := r.Run(":"+ cfg.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
