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
	tataUsahaRepo := repository.NewTataUsahaRepository(db)
	tataUsahaHandler := handlers.NewTataUsahaHandler(tataUsahaRepo, userRepo)
	siswaRepo := repository.NewSiswaRepository(db)
	pembayaranRepo := repository.NewPembayaranRepository(db)
	siswaHandler := handlers.NewSiswaHandler(siswaRepo, userRepo, pembayaranRepo)

	r := gin.Default()
	r.Use(cors.Default())

	user := r.Group("/user")
	{
		user.POST("/login", userHandler.Login)
	}
	spp := r.Group("/spp")
	{
		spp.GET("/get", sppHandler.GetSpp)
		spp.GET("/get/:id", sppHandler.GetSppById)
		spp.POST("/create", sppHandler.CreateSpp)
		spp.PUT("/update", sppHandler.UpdateSpp)
		spp.DELETE("/delete/:id", sppHandler.DeleteSpp)
	}
	kelas := r.Group("/kelas")
	{
		kelas.GET("/get", kelasHandler.GetKelas)
		kelas.GET("/get/:id", kelasHandler.GetKelasById)
		kelas.POST("/create", kelasHandler.CreateKelas)
		kelas.PUT("/update", kelasHandler.UpdateKelas)
		kelas.DELETE("/delete/:id", kelasHandler.DeleteKelas)
	}
	tataUsaha := r.Group("/tata-usaha")
	{
		tataUsaha.GET("/get", tataUsahaHandler.GetTataUsaha)
		// tataUsaha.GET("/get/:id", tataUsahaHandler.GetTataUsahaById)
		tataUsaha.POST("/create", tataUsahaHandler.CreateTataUsaha)
		tataUsaha.PUT("/update", tataUsahaHandler.UpdateTataUsaha)
		tataUsaha.DELETE("/delete/:id", tataUsahaHandler.DeleteTataUsaha)
	}
	siswa := r.Group("/siswa")
	{
		siswa.GET("/get", siswaHandler.GetSiswa)
		// tataUsaha.GET("/get/:id", tataUsahaHandler.GetTataUsahaById)
		siswa.POST("/create", siswaHandler.CreateSiswa)
		// tataUsaha.PUT("/update", tataUsahaHandler.UpdateTataUsaha)
		// tataUsaha.DELETE("/delete/:id", tataUsahaHandler.DeleteTataUsaha)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
