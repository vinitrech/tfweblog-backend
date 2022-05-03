package models

import (
	"time"

	"gorm.io/gorm"
)

type Veiculo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Modelo    string         `json:"modelo"`
	Placa     string         `json:"placa"`
	Ativo     bool           `json:"ativo"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
