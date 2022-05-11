package controllers

import (
	"strconv"
	"tfweblog/database"
	"tfweblog/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowDocumentos(c *gin.Context) {
	db := database.GetDatabase()

	var documentos []struct {
		ID            uint           `json:"id" gorm:"primaryKey"`
		Id_usuario    uint           `json:"id_usuario"`
		Id_transporte uint           `json:"id_transporte"`
		Usuario       string         `json:"usuario"`
		Titulo        string         `json:"titulo"`
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
		err := db.Raw("select documentos.*, usuarios.nome as usuario from documentos, usuarios where documentos.id_usuario = usuarios.id AND documentos.id_transporte = ? AND ((documentos.id = ?) OR (documentos.titulo ILIKE ?) OR (usuarios.nome ILIKE ?)) order by documentos.id desc", newIdTransporte, searchId, search, search).Order("id desc").Scan(&documentos).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros. " + err.Error(),
			})
			return
		}
	} else {
		err := db.Raw("select documentos.*, usuarios.nome as usuario from documentos, usuarios where documentos.id_usuario = usuarios.id and id_transporte = ? order by documentos.id desc", newIdTransporte).Order("id desc").Scan(&documentos).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros.",
			})
			return
		}
	}

	c.JSON(200, documentos)
}

func CreateDocumento(c *gin.Context) {
	db := database.GetDatabase()

	var documento models.Documento

	err := c.ShouldBindJSON(&documento)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	err = db.Create(&documento).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
