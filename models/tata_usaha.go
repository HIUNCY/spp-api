package models

type TataUsaha struct {
	IdTu     int    `json:"id_tu" gorm:"primaryKey"`
	IdUser   int    `json:"id_user"`
	NamaTu   string `json:"nama_tu"`
	NoTelpTu string `json:"no_telp_tu"`
}
