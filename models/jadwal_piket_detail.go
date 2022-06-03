package models

import "gorm.io/gorm"

type JadwalPiketDetail struct {
	gorm.Model
	Pelaksana     []AnakPanti `gorm:"many2many:jadwal_piket_detail_anak_panti;"`
	JadwalPiket   JadwalPiket
	JadwalPiketID int
	Waktu         string
}
