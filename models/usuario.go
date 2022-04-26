package models

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Cargo     string    `json:"cargo" gorm:"default:motorista"`
	Nome      string    `json:"nome"`
	Cpf       string    `json:"cpf"`
	Ativo     bool      `json:"ativo"`
	Email     string    `gorm:"unique" json:"email"`
	Senha     string    `json:"senha"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}