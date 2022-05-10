package routes

import (
	"tfweblog/controllers"
	"tfweblog/server/middlewares"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	router.Use(CORSMiddleware())

	api := router.Group("api/v1")
	{
		api.POST("login", controllers.Login)
		api.POST("cadastro", controllers.Cadastro)
		api.POST("auth/google", controllers.GoogleLogin)
		api.GET("/getData", controllers.GetData)

		api.Use(middlewares.Auth())

		api.GET("/dashboard", controllers.Dashboard)

		usuarios := api.Group("usuarios")
		{
			usuarios.GET("/", middlewares.AuthAdmin(), controllers.ShowUsuarios)
			usuarios.POST("/", middlewares.AuthAdmin(), controllers.CreateUsuario)
			usuarios.GET("/:id", middlewares.AuthAdmin(), controllers.ShowUsuario)
			usuarios.PUT("/:id", middlewares.AuthAdmin(), controllers.UpdateUsuario)
			usuarios.DELETE("/:id", middlewares.AuthAdmin(), controllers.DeleteUsuario)
			usuarios.GET("/getMotoristas", middlewares.AuthAdmin(), controllers.GetMotoristas)
			usuarios.GET("/getAdministradores", middlewares.AuthAdmin(), controllers.GetAdministradores)
			usuarios.GET("/getSupervisores", middlewares.AuthAdmin(), controllers.GetSupervisores)
		}

		veiculos := api.Group("veiculos")
		{
			veiculos.GET("/", middlewares.AuthAdminOrSupervisor(), controllers.ShowVeiculos)
			veiculos.POST("/", middlewares.AuthAdminOrSupervisor(), controllers.CreateVeiculo)
			veiculos.GET("/:id", middlewares.AuthAdminOrSupervisor(), controllers.ShowVeiculo)
			veiculos.PUT("/:id", middlewares.AuthAdminOrSupervisor(), controllers.UpdateVeiculo)
			veiculos.DELETE("/:id", middlewares.AuthAdminOrSupervisor(), controllers.DeleteVeiculo)
			veiculos.GET("/getVeiculos", middlewares.AuthAdminOrSupervisor(), controllers.GetVeiculos)
		}

		categorias := api.Group("categorias")
		{
			categorias.GET("/", middlewares.AuthAdminOrSupervisor(), controllers.ShowCategorias)
			categorias.POST("/", middlewares.AuthAdminOrSupervisor(), controllers.CreateCategoria)
			categorias.GET("/:id", middlewares.AuthAdminOrSupervisor(), controllers.ShowCategoria)
			categorias.PUT("/:id", middlewares.AuthAdminOrSupervisor(), controllers.UpdateCategoria)
			categorias.DELETE("/:id", middlewares.AuthAdminOrSupervisor(), controllers.DeleteCategoria)
		}

		clientes := api.Group("clientes")
		{
			clientes.GET("/getClientes", middlewares.AuthAdminOrSupervisor(), controllers.GetClientes)
			clientes.GET("/", middlewares.AuthAdminOrSupervisor(), controllers.ShowClientes)
			clientes.POST("/", middlewares.AuthAdminOrSupervisor(), controllers.CreateCliente)
			clientes.GET("/:id", middlewares.AuthAdminOrSupervisor(), controllers.ShowCliente)
			clientes.PUT("/:id", middlewares.AuthAdminOrSupervisor(), controllers.UpdateCliente)
			clientes.DELETE("/:id", middlewares.AuthAdminOrSupervisor(), controllers.DeleteCliente)
		}

		transportes := api.Group("transportes")
		{
			transportes.GET("/", CORSMiddleware(), controllers.ShowTransportes)
			transportes.POST("/", CORSMiddleware(), middlewares.AuthAdminOrSupervisor(), controllers.CreateTransporte)
			transportes.GET("/:id", CORSMiddleware(), middlewares.AuthAdminOrSupervisor(), controllers.ShowTransporte)
			transportes.PUT("/:id", CORSMiddleware(), middlewares.AuthAdminOrSupervisor(), controllers.UpdateTransporte)
			transportes.DELETE("/:id", CORSMiddleware(), middlewares.AuthAdminOrSupervisor(), controllers.DeleteTransporte)
			transportes.GET("/enviar-inicio-supervisor/:id", CORSMiddleware(), middlewares.AuthAdminOrDriver(), controllers.EnviarInicioSupervisor)
			transportes.GET("/aprovar-inicio/:id", CORSMiddleware(), middlewares.AuthAdminOrSupervisor(), controllers.AprovarInicio)
			transportes.GET("/enviar-finalizacao-supervisor/:id", CORSMiddleware(), middlewares.AuthAdminOrDriver(), controllers.EnviarFinalizacaoSupervisor)
			transportes.GET("/finalizar/:id", CORSMiddleware(), middlewares.AuthAdminOrSupervisor(), controllers.Finalizar)
		}

		cidades := api.Group("cidades")
		{
			cidades.GET("/", controllers.ShowCidades)
		}
	}

	return router
}
