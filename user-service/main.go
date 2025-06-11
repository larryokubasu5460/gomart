package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/larryokubasu5460/gomart/user-service/config"
	"github.com/larryokubasu5460/gomart/user-service/handler"
	"github.com/larryokubasu5460/gomart/user-service/middleware"
	"github.com/larryokubasu5460/gomart/user-service/models"
	"github.com/larryokubasu5460/gomart/user-service/repository"
	"github.com/larryokubasu5460/gomart/user-service/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()

	db, err := gorm.Open(postgres.Open(config.Cfg.DBurl), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})

	repo := &repository.UserRepository{DB: db}
	svc := &service.UserService{Repo:repo}
	h := &handler.UserHandler{Service:svc}

	r := gin.Default()
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	r.GET("/me", middleware.AuthMiddleware(), func(c *gin.Context){
		user := c.MustGet("user").(jwt.MapClaims)
		c.JSON(200, gin.H{
			"email":user["email"],
			"id": user["user_id"],
		})
	})

	r.Run(":"+config.Cfg.ServerPort)
}