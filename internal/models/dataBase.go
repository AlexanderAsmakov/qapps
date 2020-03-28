package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
)

var db *sqlx.DB

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

	db, err = sqlx.Connect("postgres", dbUri)

	if err != nil {
		log.Fatal(err)
	}

	initQuery, err := ioutil.ReadFile("db/001-init-schema.sql")

	if err != nil {
		log.Fatal(err)
	}

	dumpAppQuery, err := ioutil.ReadFile("db/002-apps-dump-data.sql")

	if err != nil {
		log.Fatal(err)
	}

	dumpBuildsQuery, err := ioutil.ReadFile("db/003-builds-dump-data.sql")

	if err != nil {
		log.Fatal(err)
	}

	dumpAppsMetadataQuery, err := ioutil.ReadFile("db/004-apps-metadata-dump-data.sql")

	if err != nil {
		log.Fatal(err)
	}

	// Initial data base structure
	db.MustExec(string(initQuery))

	// Dump test data
	tx := db.MustBegin()
	tx.MustExec(string(dumpAppQuery))
	tx.MustExec(string(dumpBuildsQuery))
	tx.MustExec(string(dumpAppsMetadataQuery))
	tx.Commit()
}

func GetDB() *sqlx.DB {
	return db
}
