package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiuncy/spp/models"
	"github.com/hiuncy/spp/repository"
)

type kelasHandler struct {
	kelasRepo repository.KelasRepo
}

func NewKelasHandler(kelasRepo repository.KelasRepo) *kelasHandler {
	return &kelasHandler{kelasRepo}
}

func (k *kelasHandler) CreateKelas(c *gin.Context) {
	var kelas models.Kelas
	if err := c.ShouldBindJSON(&kelas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := k.kelasRepo.CreateKelas(&kelas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses menambahkan kelas"})
}
