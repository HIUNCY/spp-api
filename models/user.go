package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Level    string `json:"level"`
	Gambar   string `json:"gambar"`
}
