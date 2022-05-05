package models

import (
	"time"

	"gorm.io/gorm"
)

type Transporte struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	Id_categoria     uint           `json:"id_categoria"`
	Id_cidade        uint           `json:"id_cidade"`
	Id_cliente       uint           `json:"id_cliente"`
	Id_motorista     uint           `json:"id_motorista"`
	Id_veiculo       uint           `json:"id_veiculo"`
	Id_administrador uint           `json:"id_administrador"`
	Id_supervisor    uint           `json:"id_supervisor"`
	Data_inicio      string         `json:"data_inicio"`
	Data_finalizacao string         `json:"data_finalizacao"`
	Status           string         `json:"status"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
