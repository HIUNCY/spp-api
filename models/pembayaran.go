package models

import "time"

type Pembayaran struct {
	IdPembayaran int        `json:"id_pembayaran" gorm:"primaryKey"`
	IdSiswa      int        `json:"id_siswa"`
	IdSpp        int        `json:"id_spp"`
	TglBayar     *time.Time `json:"tgl_bayar"`
	BulanDibayar string     `json:"bulan_dibayar"`
	TahunDibayar string     `json:"tahun_dibayar"`
	JumlahBayar  int        `json:"jumlah_bayar"`
}
