package controllers

import (
	"strconv"
	"tfweblog/database"
	"tfweblog/models"
	"tfweblog/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowTransportes(c *gin.Context) {
	db := database.GetDatabase()

	var arrayTransportes []struct {
		ID               uint           `json:"id" gorm:"primaryKey"`
		Id_categoria     uint           `json:"id_categoria"`
		Id_cidade        uint           `json:"id_cidade"`
		Id_cliente       uint           `json:"id_cliente"`
		Id_motorista     uint           `json:"id_motorista"`
		Id_veiculo       uint           `json:"id_veiculo"`
		Id_administrador uint           `json:"id_administrador"`
		Id_supervisor    uint           `json:"id_supervisor"`
		Data_inicio      string         `json:"data_inicio"`
		Data_finalizacao string         `json:"data_finalizacao"`
		Status           string         `json:"status"`
		CreatedAt        time.Time      `json:"created_at"`
		UpdatedAt        time.Time      `json:"updated_at"`
		DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
		Cliente          string         `json:"cliente"`
		Motorista        string         `json:"motorista"`
		Cidade           string         `json:"cidade"`
		Veiculo          string         `json:"veiculo"`
		Categoria        string         `json:"categoria"`
	}

	search, hasSearch := c.GetQuery("search")
	searchId, _ := strconv.Atoi(search)
	search = "%" + search + "%"

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

	var usuario models.Usuario

	err = db.Select([]string{"id", "nome", "email", "cargo"}).Where("id = ?", userId).Find(&usuario).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado: " + err.Error(),
		})
		return
	}

	if usuario.Cargo == "motorista" {
		if hasSearch {
			err := db.Raw(`
			select transportes.*, clientes.nome as cliente, usuarios.nome as motorista, veiculos.placa as veiculo, CONCAT(cidades.nome, ' - ',estados.sigla) as cidade, categorias.descricao as categoria
			from transportes, clientes, usuarios, veiculos, cidades, estados, categorias 
			where transportes.id_cliente = clientes.id 
			and usuarios.id = transportes.id_motorista
			and veiculos.id = transportes.id_veiculo
			and categorias.id = transportes.id_categoria
			and transportes.id_cidade = cidades.id
			and estados.id = cidades.id_estado
			and (transportes.id = ? or clientes.nome ILIKE ? or transportes.status ILIKE ?)
			and transportes.id_motorista = ?
			and transportes.deleted_at is null
			order by transportes.id desc`, searchId, search, search, usuario.ID).Scan(&arrayTransportes).Error

			if err != nil {
				c.JSON(400, gin.H{
					"error": "Não foi possível listar os registros. " + err.Error(),
				})
				return
			}
		} else {
			err := db.Raw(`select transportes.*, clientes.nome as cliente, usuarios.nome as motorista, veiculos.placa as veiculo, CONCAT(cidades.nome, ' - ',estados.sigla) as cidade, categorias.descricao as categoria
			from transportes, clientes, usuarios, veiculos, cidades, estados, categorias 
			where transportes.id_cliente = clientes.id 
			and usuarios.id = transportes.id_motorista
			and transportes.id_categoria = categorias.id
			and veiculos.id = transportes.id_veiculo
			and transportes.id_cidade = cidades.id
			and estados.id = cidades.id_estado
			and transportes.id_motorista = ?
			and transportes.deleted_at is null
			order by transportes.id desc`, usuario.ID).Scan(&arrayTransportes).Error

			if err != nil {
				c.JSON(400, gin.H{
					"error": "Não foi possível listar os registros.",
				})
				return
			}
		}
	} else {
		if hasSearch {
			err := db.Raw(`select transportes.*, clientes.nome as cliente, usuarios.nome as motorista, veiculos.placa as veiculo, CONCAT(cidades.nome, ' - ',estados.sigla) as cidade, categorias.descricao as categoria
			from transportes, clientes, usuarios, veiculos, cidades, estados, categorias 
			where transportes.id_cliente = clientes.id 
			and usuarios.id = transportes.id_motorista
			and categorias.id = transportes.id_categoria
			and veiculos.id = transportes.id_veiculo
			and transportes.id_cidade = cidades.id
			and estados.id = cidades.id_estado
			and (transportes.id = ? or clientes.nome ILIKE ? or transportes.status ILIKE ?)
			and transportes.deleted_at is null
			order by transportes.id desc`, searchId, search, search).Scan(&arrayTransportes).Error

			if err != nil {
				c.JSON(400, gin.H{
					"error": "Não foi possível listar os registros. " + err.Error(),
				})
				return
			}
		} else {
			err := db.Raw(`select transportes.*, clientes.nome as cliente, usuarios.nome as motorista, veiculos.placa as veiculo, CONCAT(cidades.nome, ' - ',estados.sigla) as cidade, categorias.descricao as categoria
			from transportes, clientes, usuarios, veiculos, cidades, estados, categorias 
			where transportes.id_cliente = clientes.id 
			and usuarios.id = transportes.id_motorista
			and categorias.id = transportes.id_categoria
			and veiculos.id = transportes.id_veiculo
			and transportes.id_cidade = cidades.id
			and estados.id = cidades.id_estado
			and transportes.deleted_at is null
			order by transportes.id desc`).Scan(&arrayTransportes).Error

			if err != nil {
				c.JSON(400, gin.H{
					"error": "Não foi possível listar os registros.",
				})
				return
			}
		}
	}

	c.JSON(200, arrayTransportes)
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

func EnviarInicioSupervisor(c *gin.Context) {

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

	if usuario.Cargo == "supervisor" {
		c.AbortWithStatus(401)
	}

	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db = database.GetDatabase()

	var transporte models.Transporte

	err = db.First(&transporte, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}

	transporte.Status = "aguardando aprovação"

	err = db.Save(&transporte).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível alterar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func AprovarInicio(c *gin.Context) {
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

	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db = database.GetDatabase()

	var transporte models.Transporte

	err = db.First(&transporte, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}

	transporte.Status = "em andamento"

	err = db.Save(&transporte).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível alterar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func EnviarFinalizacaoSupervisor(c *gin.Context) {

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

	if usuario.Cargo == "supervisor" {
		c.AbortWithStatus(401)
	}

	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db = database.GetDatabase()

	var transporte models.Transporte

	err = db.First(&transporte, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}

	transporte.Status = "aguardando finalização"

	err = db.Save(&transporte).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível alterar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func Finalizar(c *gin.Context) {

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

	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser inteiro",
		})
		return
	}

	db = database.GetDatabase()

	var transporte models.Transporte

	err = db.First(&transporte, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}

	transporte.Status = "finalizado"
	t := time.Now()
	transporte.Data_finalizacao = t.Format("2006-01-02")

	err = db.Save(&transporte).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível alterar o registro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
