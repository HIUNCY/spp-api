package repository

import (
	"github.com/hiuncy/spp/models"
	"gorm.io/gorm"
)

type KelasRepo interface {
	GetKelas() (*[]models.Kelas, error)
	GetKelasById(idKelas int) (*[]models.Kelas, error)
	CreateKelas(kelas *models.Kelas) error
	UpdateKelas(kelas *models.Kelas) error
	DeleteKelas(idSpp int) error
}

type kelasRepo struct {
	db *gorm.DB
}

func (k *kelasRepo) GetKelas() (*[]models.Kelas, error) {
	var kelas []models.Kelas
	if err := k.db.Find(&kelas).Error; err != nil {
		return nil, err
	}
	return &kelas, nil
}

func (k *kelasRepo) GetKelasById(idKelas int) (*[]models.Kelas, error) {
	var kelas []models.Kelas
	if err := k.db.First(&kelas, idKelas).Error; err != nil {
		return nil, err
	}
	return &kelas, nil
}

func (k *kelasRepo) CreateKelas(kelas *models.Kelas) error {
	if err := k.db.Create(kelas).Error; err != nil {
		return err
	}
	return nil
}

func (k *kelasRepo) UpdateKelas(kelas *models.Kelas) error {
	if err := k.db.Save(kelas).Error; err != nil {
		return err
	}
	return nil
}

func (k *kelasRepo) DeleteKelas(idKelas int) error {
	if err := k.db.Delete(&models.Kelas{}, idKelas).Error; err != nil {
		return err
	}
	return nil
}

func NewKelasRepository(db *gorm.DB) KelasRepo {
	return &kelasRepo{db}
}
