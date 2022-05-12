package controllers

import (
	"strconv"
	"tfweblog/database"
	"tfweblog/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowAvisos(c *gin.Context) {
	db := database.GetDatabase()

	var avisos []struct {
		ID            uint           `json:"id" gorm:"primaryKey"`
		Id_usuario    uint           `json:"id_usuario"`
		Id_transporte uint           `json:"id_transporte"`
		Usuario       string         `json:"usuario"`
		Descricao     string         `json:"descricao"`
		Link          string         `json:"link"`
		CreatedAt     time.Time      `json:"created_at"`
		UpdatedAt     time.Time      `json:"updated_at"`
		DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	}

	search, hasSearch := c.GetQuery("search")
	searchId, _ := strconv.Atoi(search)
	search = "%" + search + "%"

	idTransporte := c.Param("id")

	newIdTransporte, err := strconv.Atoi(idTransporte)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id do transporte precisa ser inteiro",
		})
		return
	}

	if hasSearch {
		err := db.Raw("select avisos.*, usuarios.nome as usuario from avisos, usuarios where avisos.id_usuario = usuarios.id AND avisos.id_transporte = ? AND ((avisos.id = ?) OR (avisos.descricao ILIKE ?) OR (usuarios.nome ILIKE ?)) order by avisos.id desc", newIdTransporte, searchId, search, search).Order("id desc").Scan(&avisos).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros. " + err.Error(),
			})
			return
		}
	} else {
		err := db.Raw("select avisos.*, usuarios.nome as usuario from avisos, usuarios where avisos.id_usuario = usuarios.id and id_transporte = ? order by avisos.id desc", newIdTransporte).Order("id desc").Scan(&avisos).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros.",
			})
			return
		}
	}

	c.JSON(200, avisos)
}

func CreateAviso(c *gin.Context) {
	db := database.GetDatabase()

	var aviso models.Aviso

	err := c.ShouldBindJSON(&aviso)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	err = db.Create(&aviso).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func ShowAviso(c *gin.Context) {
	id := c.Param("idAviso")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var aviso models.Aviso

	err = db.First(&aviso, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, aviso)
}
