package repository

import (
	"github.com/hiuncy/spp/models"
	"gorm.io/gorm"
)

type PembayaranRepo interface {
	CreatePembayaran(idSiswa int, idSpp int) error
}

type pembayaranRepo struct {
	db *gorm.DB
}

// CreatePembayaran implements PembayaranRepo.
func (p *pembayaranRepo) CreatePembayaran(idSiswa int, idSpp int) error {
	months := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	var pembayaranList []models.Pembayaran
	for _, month := range months {
		pembayaran := models.Pembayaran{
			IdSiswa:      idSiswa,
			IdSpp:        idSpp,
			BulanDibayar: month,
		}
		pembayaranList = append(pembayaranList, pembayaran)
	}
	if err := p.db.Create(&pembayaranList).Error; err != nil {
		return err
	}
	return nil
}

func NewPembayaranRepository(db *gorm.DB) PembayaranRepo {
	return &pembayaranRepo{db}
}
