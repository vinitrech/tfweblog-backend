package models

import (
	"time"

	"gorm.io/gorm"
)

type Estado struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Nome      string         `json:"nome"`
	Sigla     string         `json:"sigla"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
