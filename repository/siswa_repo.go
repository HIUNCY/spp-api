package repository

import (
	"github.com/hiuncy/spp/models"
	"gorm.io/gorm"
)

type SiswaRepo interface {
	GetSiswa() (*[]models.GetSiswa, error)
	// GetSiswaById(idTataUsaha int) (*models.GetDetailSiswa, error)
	CreateSiswa(idUser int, siswa *models.Siswa) (int, error)
	// UpdateSiswa(siswa *models.Siswa) error
}

type siswaRepo struct {
	db *gorm.DB
}

func (t *siswaRepo) GetSiswa() (*[]models.GetSiswa, error) {
	var results []models.GetSiswa
	if err := t.db.Model(&models.Siswa{}).Select("siswa.nisn, siswa.nama_siswa, kelas.kelas, user.gambar").Joins("inner join user ON siswa.id_user = user.id_user").Joins("inner join kelas ON siswa.id_kelas = kelas.id_kelas").Scan(&results).Error; err != nil {
		return nil, err
	}
	return &results, nil
}

// func (t *tataUsahaRepo) GetTataUsahaById(idTataUsaha int) (*[]models.TataUsaha, error) {
// 	var tataUsaha []models.TataUsaha
// 	if err := t.db.First(&tataUsaha, idTataUsaha).Error; err != nil {
// 		return nil, err
// 	}
// 	return &tataUsaha, nil
// }

func (s *siswaRepo) CreateSiswa(idUser int, siswa *models.Siswa) (int, error) {
	siswa.IdUser = idUser
	if err := s.db.Create(siswa).Error; err != nil {
		return 0, err
	}
	return siswa.IdSiswa, nil
}

// func (t *tataUsahaRepo) UpdateTataUsaha(tataUsaha *models.TataUsaha) error {
// 	if err := t.db.Save(tataUsaha).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func NewSiswaRepository(db *gorm.DB) SiswaRepo {
	return &siswaRepo{db}
}
