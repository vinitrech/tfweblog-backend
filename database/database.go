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
	// str := "postgres://armmewwrqgrkmy:8399afab3ef2da1a2603a99eb20bf5953e18b0843ce3bb13e49e29a3df300382@ec2-34-227-120-79.compute-1.amazonaws.com:5432/dcbja375s7or9g"

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