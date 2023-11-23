package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func InitializeDB() error {
	cnxn := "postgres://postgres:postgres@localhost:5432/POLONLINEGW?sslmode=disable"

	db, err := gorm.Open(postgres.Open(cnxn), &gorm.Config{})
	if err != nil {
		return err
	}

	if sqlDB, err := db.DB(); err == nil {
		if err := sqlDB.Ping(); err != nil {
			return err
		}
	} else {
		return err
	}

	gormDB = db

	log.Println("Database connection initialized successfully.")
	return nil
}

func GetGormDB() *gorm.DB {
	return gormDB
}
