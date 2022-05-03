package migrations

import (
	"tfweblog/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Usuario{})
	db.AutoMigrate(models.Veiculo{})
	db.AutoMigrate(models.Cidade{})
	db.AutoMigrate(models.Estado{})
	db.AutoMigrate(models.Cliente{})
	db.AutoMigrate(models.Categorias{})
}