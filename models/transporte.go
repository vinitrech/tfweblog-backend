package models

import (
	"time"

	"gorm.io/gorm"
)

type Transporte struct {
	ID               uint `json:"id" gorm:"primaryKey"`
	Id_categoria     uint
	Categoria        Categorias `json:"categoria" gorm:"ForeignKey:Id_categoria"`
	Id_cidade        uint
	Cidade           Cidade `json:"cidade" gorm:"ForeignKey:Id_cidade"`
	Id_cliente       uint
	Cliente          Cliente `json:"cliente" gorm:"ForeignKey:Id_cliente"`
	Id_motorista     uint
	Motorista        Usuario `json:"motorista" gorm:"ForeignKey:Id_motorista"`
	Id_veiculo       uint
	Veiculo          Veiculo `json:"veiculo" gorm:"ForeignKey:Id_veiculo"`
	Id_administrador uint
	Administrador    Usuario `json:"administrador" gorm:"ForeignKey:Id_administrador"`
	Id_supervisor    uint
	Supervisor       Usuario        `json:"supervisor" gorm:"ForeignKey:Id_supervisor"`
	Data_inicio      string         `json:"data_inicio"`
	Data_finalizacao string         `json:"data_finalizacao"`
	Status           string         `json:"status"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
