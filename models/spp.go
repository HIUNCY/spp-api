package models

type Spp struct {
	IdSpp   int `json:"id_spp" gorm:"primaryKey"`
	Tahun   int `json:"tahun"`
	Nominal int `json:"nominal"`
}
