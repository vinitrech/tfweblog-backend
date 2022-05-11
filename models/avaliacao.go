package models

import (
	"time"

	"gorm.io/gorm"
)

type Avaliacoes struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Id_usuario    uint           `json:"id_usuario"`
	Id_transporte uint           `json:"id_transporte"`
	Descricao     string         `json:"descricao"`
	Status        string         `json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
