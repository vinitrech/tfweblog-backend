package controllers

import (
	"strconv"
	"tfweblog/database"
	"tfweblog/models"

	"github.com/gin-gonic/gin"
)

func ShowVeiculos(c *gin.Context) {
	db := database.GetDatabase()

	var veiculos []models.Veiculo

	search, hasSearch := c.GetQuery("search")
	searchId, _ := strconv.Atoi(search)
	search = "%" + search + "%"

	if hasSearch {
		err := db.Where("(id = ?) OR (modelo ILIKE ?) or (placa ILIKE ?)", searchId, search, search).Order("id desc").Find(&veiculos).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros. " + err.Error(),
			})
			return
		}
	} else {
		err := db.Order("id desc").Find(&veiculos).Error

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Não foi possível listar os registros.",
			})
			return
		}
	}

	c.JSON(200, veiculos)
}

func ShowVeiculo(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var veiculo models.Veiculo

	err = db.First(&veiculo, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, veiculo)
}

func CreateVeiculo(c *gin.Context) {
	db := database.GetDatabase()

	var veiculo models.Veiculo

	err := c.ShouldBindJSON(&veiculo)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	err = db.Create(&veiculo).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func UpdateVeiculo(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var veiculo models.Veiculo

	err = db.First(&veiculo, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}

	err = c.ShouldBindJSON(&veiculo)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível converter o corpo da requisição: " + err.Error(),
		})
		return
	}

	err = db.Save(&veiculo).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível alterar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func DeleteVeiculo(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Veiculo{}, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível deletar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func GetVeiculos(c *gin.Context) {
	db := database.GetDatabase()

	var veiculos []models.Veiculo

	err := db.Where("ativo = true").Order("id desc").Find(&veiculos).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível listar os registros. " + err.Error(),
		})
		return
	}

	c.JSON(200, veiculos)
}
