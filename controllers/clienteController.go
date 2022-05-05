package controllers

import (
	"strconv"
	"tfweblog/database"
	"tfweblog/models"

	"github.com/gin-gonic/gin"
)

func ShowClientes(c *gin.Context) {
	db := database.GetDatabase()

	var clientes []models.Cliente

	search, hasSearch := c.GetQuery("search")
	searchId, _ := strconv.Atoi(search)
	search = "%" + search + "%"

	if hasSearch {
		err := db.Where("(id = ?) OR (nome ILIKE ?) or (cnpj ILIKE ?)", searchId, search, search).Order("id desc").Find(&clientes).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros. " + err.Error(),
			})
			return
		}
	} else {
		err := db.Order("id desc").Find(&clientes).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros.",
			})
			return
		}
	}

	c.JSON(200, clientes)
}

func ShowCliente(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var cliente models.Cliente

	err = db.First(&cliente, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, cliente)
}

func CreateCliente(c *gin.Context) {
	db := database.GetDatabase()

	var cliente models.Cliente

	err := c.ShouldBindJSON(&cliente)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	err = db.Create(&cliente).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func UpdateCliente(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var cliente models.Cliente

	err = db.First(&cliente, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}

	err = c.ShouldBindJSON(&cliente)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	err = db.Save(&cliente).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível alterar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func DeleteCliente(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Cliente{}, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível deletar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func GetClientes(c *gin.Context) {
	db := database.GetDatabase()

	var clientes []models.Cliente

	err := db.Where("ativo = true").Order("id desc").Find(&clientes).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível listar os registros.",
		})
		return
	}

	c.JSON(200, clientes)
}