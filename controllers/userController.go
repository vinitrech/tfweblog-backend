package controllers

import (
	"strconv"
	"tfweblog/database"
	"tfweblog/models"
	"tfweblog/services"

	"github.com/gin-gonic/gin"
)

func ShowUsuarios(c *gin.Context) {
	db := database.GetDatabase()

	var usuarios []models.Usuario

	err:= db.Order("id desc").Find(&usuarios).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível listar os registros.",
		})
		return
	}

	c.JSON(200, usuarios)
}

func ShowUsuario(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var usuario models.Usuario

	err = db.First(&usuario, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, usuario)
}

func CreateUsuario(c *gin.Context) {
	db := database.GetDatabase()

	var usuario models.Usuario

	err := c.ShouldBindJSON(&usuario)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	usuario.Senha = services.SHA256Encoder(usuario.Senha)

	err = db.Create(&usuario).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func UpdateUsuario(c *gin.Context) {
	db := database.GetDatabase()

	var usuario models.Usuario

	err := c.ShouldBindJSON(&usuario)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}
	
	usuario.Senha = services.SHA256Encoder(usuario.Senha)

	err = db.Save(&usuario).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível alterar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func DeleteUsuario(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Usuario{}, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível deletar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}