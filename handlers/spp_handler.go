package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiuncy/spp/models"
	"github.com/hiuncy/spp/repository"
)

type sppHandler struct {
	sppRepo repository.SppRepo
}

func NewSppHandler(sppRepo repository.SppRepo) *sppHandler {
	return &sppHandler{sppRepo}
}

func (s *sppHandler) CreateSpp(c *gin.Context) {
	var spp models.Spp
	if err := c.ShouldBindJSON(&spp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.sppRepo.CreateSpp(&spp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses menambahkan SPP"})
}

func (s *sppHandler) UpdateSpp(c *gin.Context) {
	var spp models.Spp
	if err := c.ShouldBindJSON(&spp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.sppRepo.UpdateSpp(&spp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses memperbarui SPP"})
}

func (s *sppHandler) DeleteSpp(c *gin.Context) {
	idSpp := c.Param("id")
	idSppInt, err := strconv.Atoi(idSpp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := s.sppRepo.DeleteSpp(idSppInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sukses menghapus SPP"})
}
