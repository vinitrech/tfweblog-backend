package controllers

import (
	"strconv"
	"tfweblog/database"
	"tfweblog/models"

	"github.com/gin-gonic/gin"
)

func ShowTransportes(c *gin.Context) {
	db := database.GetDatabase()

	var transportes []models.Transporte

	search, hasSearch := c.GetQuery("search")
	searchId, _ := strconv.Atoi(search)
	search = "%" + search + "%"

	if hasSearch {
		err := db.Where("(id = ?) OR (status ILIKE ?)", searchId, search).Order("id desc").Find(&transportes).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros. " + err.Error(),
			})
			return
		}
	} else {
		err := db.Order("id desc").Find(&transportes).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros.",
			})
			return
		}
	}

	c.JSON(200, transportes)
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