package models

import "gorm.io/gorm"

type AnakPanti struct {
	gorm.Model
	UserID  int
	KamarID int
	User    User  `gorm:"unique"`
	Kamar   Kamar `gorm:"unique"`
}

type AnakPantiCreateInput struct {
	UserID  int `json:"user_id" binding:"required"`
	KamarID int `json:"kamar_id" binding:"required"`
}
