package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hiuncy/spp/handlers"
	"github.com/hiuncy/spp/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/spponline?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // ganti dengan URL React kamu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	user := r.Group("/user")
	{
		user.POST("/create", userHandler.CreateUser)
		user.POST("/login", userHandler.Login)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
