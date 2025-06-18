package models

type User struct {
	IdUser   int    `json:"id_user" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Level    string `json:"level"`
	Gambar   string `json:"gambar"`
}
