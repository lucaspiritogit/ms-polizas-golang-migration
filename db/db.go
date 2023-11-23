package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
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

func InitializeSQLSERVERDB() error {
	server := "172.16.2.41"
	port := 1433
	user := "usrFusap"
	password := "SaT2oLO!Zt2T"
	database := "POLONLINEGW"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

	var err error
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(15)
	db.SetConnMaxLifetime(time.Minute + 10)

	err = db.Ping()
	if err != nil {
		return err
	}

	gormDB, err = gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Database connection initialized successfully.")
	return nil
}

func GetGormDB() *gorm.DB {
	return gormDB
}
