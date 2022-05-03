package models

import (
	"time"

	"gorm.io/gorm"
)

type Categoria struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Descricao string         `json:"descricao"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
