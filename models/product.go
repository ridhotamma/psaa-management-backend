package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	Categories []Category `gorm:"many2many:product_categories;"`
}

type Category struct {
	gorm.Model
	Name string
}
