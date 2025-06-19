package models

type Siswa struct {
	IdSiswa   int    `json:"id_siswa" gorm:"primaryKey"`
	IdUser    int    `json:"id_user"`
	IdSpp     int    `json:"id_spp"`
	IdKelas   int    `json:"id_kelas"`
	Nisn      string `json:"nisn"`
	NamaSiswa string `json:"nama_siswa"`
	Alamat    string `json:"alamat"`
	NoTelp    string `json:"no_telp"`
}

type GetSiswa struct {
	Nisn      string `json:"nisn"`
	NamaSiswa string `json:"nama_siswa"`
	Kelas     string `json:"kelas"`
	Gambar    string `json:"gambar"`
}

type GetDetailSiswa struct{}
