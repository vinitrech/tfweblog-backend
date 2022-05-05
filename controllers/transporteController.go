package controllers

import (
	"strconv"
	"tfweblog/database"
	"tfweblog/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowTransportes(c *gin.Context) {
	db := database.GetDatabase()

	var arrayTransportes []struct {
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
		Cliente          string         `json:"cliente"`
	}

	search, hasSearch := c.GetQuery("search")
	searchId, _ := strconv.Atoi(search)
	search = "%" + search + "%"

	if hasSearch {
		err := db.Raw("select transportes.*, clientes.nome as cliente from transportes, clientes where transportes.id_cliente = clientes.id and (transportes.id = ? or clientes.nome ILIKE ? or transportes.status ILIKE ?)", searchId, search, search).Scan(&arrayTransportes).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros. " + err.Error(),
			})
			return
		}
	} else {
		err := db.Raw("select transportes.*, clientes.nome as cliente from transportes, clientes where transportes.id_cliente = clientes.id").Scan(&arrayTransportes).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros.",
			})
			return
		}
	}

	c.JSON(200, arrayTransportes)
}

func ShowTransporte(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var transporte models.Transporte

	err = db.First(&transporte, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, transporte)
}

func CreateTransporte(c *gin.Context) {
	db := database.GetDatabase()

	var transporte models.Transporte

	err := c.ShouldBindJSON(&transporte)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	err = db.Create(&transporte).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func UpdateTransporte(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var transporte models.Transporte

	err = db.First(&transporte, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}

	err = c.ShouldBindJSON(&transporte)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	err = db.Save(&transporte).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível alterar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func DeleteTransporte(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Transporte{}, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível deletar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
