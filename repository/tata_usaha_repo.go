package repository

import (
	"github.com/hiuncy/spp/models"
	"gorm.io/gorm"
)

type TataUsahaRepo interface {
	GetTataUsaha() (*[]models.GetTataUsaha, error)
	// GetTataUsahaById(idTataUsaha int) (*models.TataUsaha, error)
	CreateTataUsaha(idUser int, tataUsaha *models.TataUsaha) error
	UpdateTataUsaha(tataUsaha *models.UpdateTataUsaha) error
}

type tataUsahaRepo struct {
	db *gorm.DB
}

func (t *tataUsahaRepo) GetTataUsaha() (*[]models.GetTataUsaha, error) {
	var results []models.GetTataUsaha
	if err := t.db.Model(&models.TataUsaha{}).Select("user.id_user, user.email, user.gambar, tata_usaha.nama_tu, tata_usaha.no_telp_tu").Joins("inner join user on tata_usaha.id_user = user.id_user").Scan(&results).Error; err != nil {
		return nil, err
	}
	return &results, nil
}

// func (t *tataUsahaRepo) GetTataUsahaById(idTataUsaha int) (*models.TataUsaha, error) {
// 	var tataUsaha []models.TataUsaha
// 	if err := t.db.First(&tataUsaha, idTataUsaha).Error; err != nil {
// 		return nil, err
// 	}
// 	return &tataUsaha, nil
// }

func (t *tataUsahaRepo) CreateTataUsaha(idUser int, tataUsaha *models.TataUsaha) error {
	tataUsaha.IdUser = idUser
	if err := t.db.Create(tataUsaha).Error; err != nil {
		return err
	}
	return nil
}

func (t *tataUsahaRepo) UpdateTataUsaha(tataUsaha *models.UpdateTataUsaha) error {
	if err := t.db.Table("tata_usaha").Save(tataUsaha).Error; err != nil {
		return err
	}
	return nil
}

func NewTataUsahaRepository(db *gorm.DB) TataUsahaRepo {
	return &tataUsahaRepo{db}
}
