package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiuncy/spp/models"
	"github.com/hiuncy/spp/repository"
)

type tataUsahaHandler struct {
	tataUsahaRepo repository.TataUsahaRepo
	userRepo      repository.UserRepo
}

func NewTataUsahaHandler(tataUsahaRepo repository.TataUsahaRepo, userRepo repository.UserRepo) *tataUsahaHandler {
	return &tataUsahaHandler{tataUsahaRepo, userRepo}
}

func (t *tataUsahaHandler) GetTataUsaha(c *gin.Context) {
	tataUsaha, err := t.tataUsahaRepo.GetTataUsaha()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tataUsaha)
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

func (t *tataUsahaHandler) CreateTataUsaha(c *gin.Context) {
	var tuCreate models.TataUsahaUser
	if err := c.ShouldBindJSON(&tuCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := t.userRepo.GetUserByEmail(tuCreate.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah digunakan"})
		return
	}
	user := models.User{
		Email:    tuCreate.Email,
		Password: tuCreate.Password,
		Level:    "tu",
		Gambar:   tuCreate.Gambar,
	}
	idUser, err := t.userRepo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tataUsaha := models.TataUsaha{
		IdUser:   idUser,
		NamaTu:   tuCreate.NamaTu,
		NoTelpTu: tuCreate.NoTelpTu,
	}
	if err := t.tataUsahaRepo.CreateTataUsaha(idUser, &tataUsaha); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses menambahkan tata usaha"})
}

func (t *tataUsahaHandler) UpdateTataUsaha(c *gin.Context) {
	var tataUsaha models.TataUsaha
	if err := c.ShouldBindJSON(&tataUsaha); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := t.tataUsahaRepo.UpdateTataUsaha(&tataUsaha); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses memperbarui tata usaha"})
}

func (t *tataUsahaHandler) DeleteTataUsaha(c *gin.Context) {
	idUser := c.Param("id")
	idUserInt, err := strconv.Atoi(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := t.userRepo.DeleteUser(idUserInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses menghapus tata usaha"})
}
