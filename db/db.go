package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func InitializeDB() error {
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