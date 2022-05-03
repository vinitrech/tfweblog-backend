package models

import (
	"time"

	"gorm.io/gorm"
)

type Cidade struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Id_estado uint
	Estado    Estado         `json:"estado" gorm:"ForeignKey:Id_estado"`
	Nome      string         `json:"nome"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
