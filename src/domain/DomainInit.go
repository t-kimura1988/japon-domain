package domain

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var dbErr error

func Init() {
	dns := ""
	DB, dbErr = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("database cannot connect: %v\n", dns)
	}
}
