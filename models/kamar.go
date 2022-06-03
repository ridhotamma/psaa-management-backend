package models

import "gorm.io/gorm"

type Kamar struct {
	gorm.Model
	Name     string `gorm:"type:varchar;not null;unique"`
	Capacity int
	Members  []AnakPanti
}

type KamarCreateInput struct {
	Name     string `json:"name" binding:"required"`
	Capacity int    `json:"capacity" binding:"required"`
}

type KamarUpdateInput struct {
	Name     string `json:"name" binding:"required"`
	Capacity int    `json:"capacity" binding:"required"`
	Members  []int  `json:"members" binding:"required"`
}
