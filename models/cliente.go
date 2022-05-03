package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nome      string    `json:"nome"`
	Cnpj       string    `json:"cnpj"`
	Ativo     bool      `json:"ativo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}