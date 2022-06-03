package models

import "gorm.io/gorm"

type JadwalPiket struct {
	gorm.Model
	Name         string `gorm:"type:varchar;not null;unique"`
	JadwalDetail []JadwalPiketDetail
	Description  string
	Day          string
}
