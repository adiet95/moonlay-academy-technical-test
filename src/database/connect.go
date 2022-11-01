package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASS")
	dbName := os.Getenv("DB")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)

	gormDB, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed connection to database")
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, errors.New("failed conection to database")
	}
	db.SetConnMaxIdleTime(100)
	db.SetMaxOpenConns(10)

	return gormDB, nil
}
