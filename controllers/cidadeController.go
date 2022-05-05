package controllers

import (
	"tfweblog/database"

	"github.com/gin-gonic/gin"
)

func ShowCidades(c *gin.Context) {
	db := database.GetDatabase()

	var cidades []struct {
		ID   uint
		Nome string
	}

	err := db.Raw("select cidades.id, CONCAT(cidades.nome, ' - ', estados.sigla) as nome from cidades join estados on estados.id = cidades.id_estado order by estados.sigla, cidades.nome").Scan(&cidades).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível listar os registros. " + err.Error(),
		})
		return
	}

	c.JSON(200, cidades)
}
