package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hiuncy/spp/handlers"
	"github.com/hiuncy/spp/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/spponline?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)
	sppRepo := repository.NewSppRepository(db)
	sppHandler := handlers.NewSppHandler(sppRepo)
	kelasRepo := repository.NewKelasRepository(db)
	kelasHandler := handlers.NewKelasHandler(kelasRepo)

	r := gin.Default()
	r.Use(cors.Default())

	user := r.Group("/user")
	{
		user.POST("/login", userHandler.Login)
	}
	spp := r.Group("/spp")
	{
		spp.POST("/create", sppHandler.CreateSpp)
		spp.POST("/update", sppHandler.UpdateSpp)
		spp.GET("/delete/:id", sppHandler.DeleteSpp)
	}
	kelas := r.Group("/kelas")
	{
		kelas.POST("/create", kelasHandler.CreateKelas)
		spp.POST("/update", kelasHandler.UpdateKelas)
		spp.GET("/delete/:id", kelasHandler.DeleteKelas)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
