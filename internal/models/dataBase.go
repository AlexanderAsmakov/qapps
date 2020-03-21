package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	var err error

	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	sslMode := os.Getenv("SSL_MODE")

	dbUri := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		dbHost, username, dbName, sslMode, password)

	db, err = gorm.Open("postgres", dbUri)

	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
