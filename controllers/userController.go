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

	search, hasSearch := c.GetQuery("search")
	searchId, _ := strconv.Atoi(search)
	search = "%" + search + "%"

	if hasSearch {
		err := db.Where("(id = ?) OR (nome ILIKE ?) or (email ILIKE ?) or (cpf ILIKE ?)", searchId, search, search, search).Order("id desc").Find(&usuarios).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros. " + err.Error(),
			})
			return
		}
	} else {
		err := db.Order("id desc").Find(&usuarios).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros.",
			})
			return
		}
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
			"error": "Registro não encontrado",
		})
		return
	}

	senha := usuario.Senha

	err = c.ShouldBindJSON(&usuario)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	if len(usuario.Senha) > 0 {
		usuario.Senha = services.SHA256Encoder(usuario.Senha)
	} else {
		usuario.Senha = senha
	}

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

func GetData(c *gin.Context){
	const Bearer_schema = "Bearer "
	header := c.GetHeader("Authorization")

	token := header[len(Bearer_schema):]

	userId, err := services.NewJWTService().GetIDFromToken(token)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Não foi possível decodificar o token: " + err.Error(),
		})
		return
	}

	db := database.GetDatabase()

	var usuario models.Usuario

	err = db.Select([]string{"id", "nome", "email", "cargo"}).Where("id = ?", userId).Find(&usuario).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, usuario)
}