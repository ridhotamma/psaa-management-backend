package models

import "gorm.io/gorm"

type Pengasuh struct {
	gorm.Model
	UserID  int
	KamarID int
	User    User
	Kamar   Kamar
}

type PengasuhCreateInput struct {
	UserID  int `json:"user_id" binding:"required"`
	KamarID int `json:"kamar_id" binding:"required"`
}
