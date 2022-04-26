package controllers

import (
	"tfweblog/database"
	"tfweblog/models"
	"tfweblog/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var login models.Login

	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	var usuario models.Usuario

	dbError := db.Where("email = ?", login.Email).First(&usuario).Error

	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	if usuario.Senha != services.SHA256Encoder(login.Senha) {
		c.JSON(401, gin.H{
			"error": "Credenciais inválidas",
		})
		return
	}

	if !usuario.Ativo {
		c.JSON(401, gin.H{
			"error": "Usuário bloqueado",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(usuario.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func Cadastro(c *gin.Context) {
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

func GoogleLogin(c *gin.Context) {
	db := database.GetDatabase()

	var login models.GoogleLogin

	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	var usuario models.Usuario

	dbError := db.Where("email = ?", login.Email).First(&usuario).Error

	if dbError != nil {

		usuario.Email = login.Email

		err = db.Create(&usuario).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível criar o registro: " + err.Error(),
			})
			return
		}

		c.JSON(401, gin.H{
			"error": "Usuário bloqueado",
		})

		return
	}

	if !usuario.Ativo {
		c.JSON(401, gin.H{
			"error": "Usuário bloqueado",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(usuario.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
