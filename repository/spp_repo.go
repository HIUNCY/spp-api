package repository

import (
	"github.com/hiuncy/spp/models"
	"gorm.io/gorm"
)

type SppRepo interface {
	GetSpp() (*[]models.Spp, error)
	GetSppById(idSpp int) (*models.Spp, error)
	CreateSpp(spp *models.Spp) error
	UpdateSpp(spp *models.Spp) error
	DeleteSpp(idSpp int) error
}

type sppRepo struct {
	db *gorm.DB
}

func (s *sppRepo) GetSpp() (*[]models.Spp, error) {
	var spp []models.Spp
	if err := s.db.Find(&spp).Error; err != nil {
		return nil, err
	}
	return &spp, nil
}

func (s *sppRepo) GetSppById(idSpp int) (*models.Spp, error) {
	var spp models.Spp
	if err := s.db.First(&spp, idSpp).Error; err != nil {
		return nil, err
	}
	return &spp, nil
}

func (s *sppRepo) CreateSpp(spp *models.Spp) error {
	if err := s.db.Create(spp).Error; err != nil {
		return err
	}
	return nil
}

func (s *sppRepo) UpdateSpp(spp *models.Spp) error {
	if err := s.db.Save(spp).Error; err != nil {
		return err
	}
	return nil
}

func (s *sppRepo) DeleteSpp(idSpp int) error {
	if err := s.db.Delete(&models.Spp{}, idSpp).Error; err != nil {
		return err
	}
	return nil
}

func NewSppRepository(db *gorm.DB) SppRepo {
	return &sppRepo{db}
}
