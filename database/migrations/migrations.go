package migrations

import (
	"tfweblog/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Usuario{})
}