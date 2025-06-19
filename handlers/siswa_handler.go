package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiuncy/spp/models"
	"github.com/hiuncy/spp/repository"
)

type siswaHandler struct {
	siswaRepo      repository.SiswaRepo
	userRepo       repository.UserRepo
	pembayaranRepo repository.PembayaranRepo
}

func NewSiswaHandler(siswaRepo repository.SiswaRepo, userRepo repository.UserRepo, pembayaranRepo repository.PembayaranRepo) *siswaHandler {
	return &siswaHandler{siswaRepo, userRepo, pembayaranRepo}
}

func (s *siswaHandler) GetSiswa(c *gin.Context) {
	siswa, err := s.siswaRepo.GetSiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, siswa)
}

// func (t *tataUsahaHandler) GetTataUsahaById(c *gin.Context) {
// 	idTataUsaha := c.Param("id")
// 	idTataUsahaInt, err := strconv.Atoi(idTataUsaha)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
// 		return
// 	}
// 	kelas, err := t.tataUsahaRepo.GetTataUsahaById(idTataUsahaInt)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, kelas)
// }

func (s *siswaHandler) CreateSiswa(c *gin.Context) {
	var siswa models.Siswa
	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := s.userRepo.GetUserByEmail(fmt.Sprintf("%s@gmail.com", siswa.Nisn))
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah digunakan"})
		return
	}
	user := models.User{
		Email:    fmt.Sprintf("%s@gmail.com", siswa.Nisn),
		Password: siswa.Nisn,
		Level:    "siswa",
		Gambar:   "default.jpg",
	}
	idUser, err := s.userRepo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	idSiswa, err := s.siswaRepo.CreateSiswa(idUser, &siswa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := s.pembayaranRepo.CreatePembayaran(idSiswa, siswa.IdSpp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses menambahkan siswa"})
}

func (s *siswaHandler) UpdateSiswa(c *gin.Context) {
	var siswa models.Siswa
	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.siswaRepo.UpdateSiswa(&siswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses memperbarui data siswa"})
}

func (s *siswaHandler) DeleteSiswa(c *gin.Context) {
	idUser := c.Param("id")
	idUserInt, err := strconv.Atoi(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := s.userRepo.DeleteUser(idUserInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses menghapus tata usaha"})
}
