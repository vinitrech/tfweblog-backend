package models

import (
	"time"

	"gorm.io/gorm"
)

type Incidente struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Id_usuario    uint           `json:"id_usuario"`
	Id_transporte uint           `json:"id_transporte"`
	Descricao     string         `gorm:"type:text" json:"descricao"`
	Link          string         `json:"link"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
