package repository

import (
	"github.com/hiuncy/spp/models"
	"gorm.io/gorm"
)

type KelasRepo interface {
	CreateKelas(kelas *models.Kelas) error
}

type kelasRepo struct {
	db *gorm.DB
}

func (k *kelasRepo) CreateKelas(kelas *models.Kelas) error {
	if err := k.db.Create(kelas).Error; err != nil {
		return err
	}
	return nil
}

func NewKelasRepository(db *gorm.DB) KelasRepo {
	return &kelasRepo{db}
}
