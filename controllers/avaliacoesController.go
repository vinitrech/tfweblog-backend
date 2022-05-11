package controllers

import (
	"strconv"
	"tfweblog/database"
	"tfweblog/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowAvaliacoes(c *gin.Context) {
	db := database.GetDatabase()

	var avaliacoes []struct {
		ID            uint           `json:"id" gorm:"primaryKey"`
		Id_usuario    uint           `json:"id_usuario"`
		Id_transporte uint           `json:"id_transporte"`
		Usuario       string         `json:"usuario"`
		Descricao     string         `json:"descricao"`
		Status        string         `json:"status"`
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
		err := db.Raw("select avaliacoes.*, usuarios.nome as usuario from avaliacoes, usuarios where avaliacoes.id_usuario = usuarios.id AND avaliacoes.id_transporte = ? AND ((avaliacoes.id = ?) OR (avaliacoes.descricao ILIKE ?) OR (usuarios.nome ILIKE ?)) order by avaliacoes.id desc", newIdTransporte, searchId, search, search).Order("id desc").Scan(&avaliacoes).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros. " + err.Error(),
			})
			return
		}
	} else {
		err := db.Raw("select avaliacoes.*, usuarios.nome as usuario from avaliacoes, usuarios where avaliacoes.id_usuario = usuarios.id and id_transporte = ? order by avaliacoes.id desc", newIdTransporte).Order("id desc").Scan(&avaliacoes).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros.",
			})
			return
		}
	}

	c.JSON(200, avaliacoes)
}

func CreateAvaliacao(c *gin.Context) {
	db := database.GetDatabase()

	var avaliacao models.Avaliacoes

	err := c.ShouldBindJSON(&avaliacao)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	err = db.Create(&avaliacao).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func ShowAvaliacao(c *gin.Context) {
	id := c.Param("idAvaliacao")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var avaliacao models.Avaliacoes

	err = db.First(&avaliacao, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, avaliacao)
}
