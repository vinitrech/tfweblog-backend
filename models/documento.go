package models

import (
	"time"

	"gorm.io/gorm"
)

type Documento struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Id_usuario    uint           `json:"id_usuario"`
	Id_transporte uint           `json:"id_transporte"`
	Titulo        string         `json:"titulo"`
	Link          string         `json:"link"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
