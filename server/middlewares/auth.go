package middlewares

import (
	"tfweblog/database"
	"tfweblog/models"
	"tfweblog/services"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "

		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "

		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}

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

		if usuario.Cargo != "administrador" {
			c.AbortWithStatus(401)
		}
	}
}

func AuthAdminOrSupervisor() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "

		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}

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

		if usuario.Cargo == "motorista" {
			c.AbortWithStatus(401)
		}
	}
}
