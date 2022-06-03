package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `gorm:"type:varchar;not null"`
	LastName    string `gorm:"type:varchar;not null"`
	Email       string `gorm:"type:varchar;not null;unique"`
	Username    string `gorm:"type:varchar;not null;unique"`
	Password    string `gorm:"type:varchar;not null"`
	PhoneNumber string `gorm:"type:varchar;not null;unique"`
	Role        int    `gorm:"default:1;not null"`
	IsActive    bool   `gorm:"type:bool;default:true;"`
}

type UserInput struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
