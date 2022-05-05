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

	search, hasSearch := c.GetQuery("search")
	search = "%" + search + "%"

	if hasSearch {
		err := db.Raw("select cidades.id, CONCAT(cidades.nome, ' - ', estados.sigla) as nome from cidades join estados on estados.id = cidades.id_estado where cidades.nome ILIKE ? order by estados.sigla, cidades.nome limit 10", search).Scan(&cidades).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros. " + err.Error(),
			})
			return
		}
	} else {
		err := db.Raw("select cidades.id, CONCAT(cidades.nome, ' - ', estados.sigla) as nome from cidades join estados on estados.id = cidades.id_estado order by estados.sigla, cidades.nome limit 10").Scan(&cidades).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros.",
			})
			return
		}
	}

	c.JSON(200, cidades)
}
