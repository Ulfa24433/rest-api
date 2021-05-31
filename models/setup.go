package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupModels() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connect to DB")
	}

	db.AutoMigrate(&Mahasiswa{})

	return db
}
