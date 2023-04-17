package domain

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var Tran *sql.DB
var dbErr error
var sqlErr error

func Init() {
	password := os.Getenv("JAPON_DB_PASSWORD")
	host := os.Getenv("JAPON_DB_HOST")
	dbName := os.Getenv("JAPON_DB_NAME")
	username := os.Getenv("JAPON_DB_USERNAME")
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo", host, username, password, dbName)
	DB, dbErr = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("database cannot connect: %v\n", dns)
	}

	Tran, sqlErr = DB.DB()
	if sqlErr != nil {
		log.Fatalf("database cannot connect: %v\n", dns)
	}
}
