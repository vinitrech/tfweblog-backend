package database

import (
	"log"
	"os"
	"tfweblog/database/migrations"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	str := os.Getenv("DATABASE_URL")

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}