package models

type Kelas struct {
	IdKelas int `json:"id_kelas" gorm:"primaryKey"`
	Kelas   int `json:"kelas"`
}
