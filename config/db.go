package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func SetupDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=ridho2002 dbname=psaa_dev port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "psaa_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}
