package models

type Spp struct {
	IdSpp   int `json:"id_spp" gorm:"primaryKey"`
	IdKelas int `json:"id_kelas"`
	Tahun   int `json:"tahun"`
	Nominal int `json:"nominal"`
}
