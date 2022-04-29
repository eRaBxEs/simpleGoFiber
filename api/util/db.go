package util

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
