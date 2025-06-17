package repository

import (
	"github.com/hiuncy/spp/models"
	"gorm.io/gorm"
)

type TataUsahaRepo interface {
	GetTataUsaha() (*[]models.TataUsaha, error)
	GetTataUsahaById(idTataUsaha int) (*[]models.TataUsaha, error)
	CreateTataUsaha(tataUsaha *models.TataUsaha) error
	UpdateTataUsaha(tataUsaha *models.TataUsaha) error
	DeleteTataUsaha(idTataUsaha int) error
}

type tataUsahaRepo struct {
	db *gorm.DB
}

func (t *tataUsahaRepo) GetTataUsaha() (*[]models.TataUsaha, error) {
	var tataUsaha []models.TataUsaha
	if err := t.db.Find(&tataUsaha).Error; err != nil {
		return nil, err
	}
	return &tataUsaha, nil
}

func (t *tataUsahaRepo) GetTataUsahaById(idTataUsaha int) (*[]models.TataUsaha, error) {
	var tataUsaha []models.TataUsaha
	if err := t.db.First(&tataUsaha, idTataUsaha).Error; err != nil {
		return nil, err
	}
	return &tataUsaha, nil
}

func (t *tataUsahaRepo) CreateTataUsaha(tataUsaha *models.TataUsaha) error {
	if err := t.db.Create(tataUsaha).Error; err != nil {
		return err
	}
	return nil
}

func (t *tataUsahaRepo) UpdateTataUsaha(tataUsaha *models.TataUsaha) error {
	if err := t.db.Save(tataUsaha).Error; err != nil {
		return err
	}
	return nil
}

func (t *tataUsahaRepo) DeleteTataUsaha(idTataUsaha int) error {
	if err := t.db.Delete(&models.TataUsaha{}, idTataUsaha).Error; err != nil {
		return err
	}
	return nil
}

func NewTataUsahaRepository(db *gorm.DB) TataUsahaRepo {
	return &tataUsahaRepo{db}
}
