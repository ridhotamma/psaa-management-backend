package main

import (
	"psaa-api/config"
	"psaa-api/models"
	"psaa-api/routes"
)

func main() {
	db := config.SetupDB()
	db.AutoMigrate(
		&models.User{},
		&models.Pengasuh{},
		&models.AnakPanti{},
		&models.Kamar{},
		&models.JadwalPiket{},
		&models.JadwalPiketDetail{},
		&models.Product{},
	)

	r := routes.SetupRoutes(db)
	r.Run()
}
