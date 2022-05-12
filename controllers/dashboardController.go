package controllers

import (
	"tfweblog/database"

	"github.com/gin-gonic/gin"
)

func GetTransportesPorMotorista(c *gin.Context) {

	db := database.GetDatabase()

	var transportes []struct {
		Quantidade uint   `json:"quantidade"`
		Motorista  string `json:"motorista"`
	}

	err := db.Raw(`select count(*) as quantidade, id_motorista as motorista from transportes group by id_motorista order by quantidade asc`).Scan(&transportes).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, transportes)
}

func GetTransportesPorCliente(c *gin.Context) {

	db := database.GetDatabase()

	var transportes []struct {
		Quantidade uint   `json:"quantidade"`
		Cliente    string `json:"cliente"`
	}

	err := db.Raw(`select count(*) as quantidade, id_cliente as cliente from transportes group by id_cliente order by quantidade asc`).Scan(&transportes).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, transportes)
}

func GetTransportesStatus(c *gin.Context) {

	db := database.GetDatabase()

	var transportes []struct {
		Quantidade uint   `json:"quantidade"`
		Status     string `json:"status"`
	}

	err := db.Raw(`select count(*) as quantidade, status from transportes group by status order by quantidade asc`).Scan(&transportes).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, transportes)
}

func GetIncidentes(c *gin.Context) {

	db := database.GetDatabase()

	var incidentes []struct {
		Id_transporte uint `json:"id_transporte"`
		Incidentes    uint `json:"incidentes"`
	}

	err := db.Raw(`select id_transporte, count(*) as "incidentes" from incidentes group by id_transporte order by incidentes asc limit 15`).Scan(&incidentes).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, incidentes)
}
