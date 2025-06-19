package models

type TataUsaha struct {
	IdTu     int    `json:"id_tu" gorm:"primaryKey"`
	IdUser   int    `json:"id_user"`
	NamaTu   string `json:"nama_tu"`
	NoTelpTu string `json:"no_telp_tu"`
}

type TataUsahaUser struct {
	IdTu     int    `json:"id_tu" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gambar   string `json:"gambar"`
	IdUser   int    `json:"id_user"`
	NamaTu   string `json:"nama_tu"`
	NoTelpTu string `json:"no_telp_tu"`
}

type GetTataUsaha struct {
	IdUser   int    `json:"id_user"`
	Email    string `json:"email"`
	Gambar   string `json:"gambar"`
	NamaTu   string `json:"nama_tu"`
	NoTelpTu string `json:"no_telp_tu"`
}
