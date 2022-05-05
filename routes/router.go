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

		api.Use(middlewares.Auth()).GET("/dashboard", controllers.Dashboard)

		usuarios := api.Group("usuarios")
		{
			usuarios.Use(middlewares.AuthAdmin()).GET("/", controllers.ShowUsuarios)
			usuarios.Use(middlewares.AuthAdmin()).POST("/", controllers.CreateUsuario)
			usuarios.Use(middlewares.AuthAdmin()).GET("/:id", controllers.ShowUsuario)
			usuarios.Use(middlewares.AuthAdmin()).PUT("/:id", controllers.UpdateUsuario)
			usuarios.Use(middlewares.AuthAdmin()).DELETE("/:id", controllers.DeleteUsuario)
			usuarios.Use(middlewares.AuthAdminOrSupervisor()).GET("/getMotoristas", controllers.GetMotoristas)
			usuarios.Use(middlewares.AuthAdminOrSupervisor()).GET("/getAdministradores", controllers.GetAdministradores)
			usuarios.Use(middlewares.AuthAdminOrSupervisor()).GET("/getSupervisores", controllers.GetSupervisores)
		}

		veiculos := api.Group("veiculos")
		{
			veiculos.Use(middlewares.AuthAdminOrSupervisor()).GET("/", controllers.ShowVeiculos)
			veiculos.Use(middlewares.AuthAdminOrSupervisor()).POST("/", controllers.CreateVeiculo)
			veiculos.Use(middlewares.AuthAdminOrSupervisor()).GET("/:id", controllers.ShowVeiculo)
			veiculos.Use(middlewares.AuthAdminOrSupervisor()).PUT("/:id", controllers.UpdateVeiculo)
			veiculos.Use(middlewares.AuthAdminOrSupervisor()).DELETE("/:id", controllers.DeleteVeiculo)
			veiculos.Use(middlewares.AuthAdminOrSupervisor()).GET("/getVeiculos", controllers.GetVeiculos)
		}

		categorias := api.Group("categorias")
		{
			categorias.Use(middlewares.AuthAdminOrSupervisor()).GET("/", controllers.ShowCategorias)
			categorias.Use(middlewares.AuthAdminOrSupervisor()).POST("/", controllers.CreateCategoria)
			categorias.Use(middlewares.AuthAdminOrSupervisor()).GET("/:id", controllers.ShowCategoria)
			categorias.Use(middlewares.AuthAdminOrSupervisor()).PUT("/:id", controllers.UpdateCategoria)
			categorias.Use(middlewares.AuthAdminOrSupervisor()).DELETE("/:id", controllers.DeleteCategoria)
		}

		clientes := api.Group("clientes")
		{
			clientes.Use(middlewares.AuthAdminOrSupervisor()).GET("/getClientes", controllers.GetClientes)
			clientes.Use(middlewares.AuthAdminOrSupervisor()).GET("/", controllers.ShowClientes)
			clientes.Use(middlewares.AuthAdminOrSupervisor()).POST("/", controllers.CreateCliente)
			clientes.Use(middlewares.AuthAdminOrSupervisor()).GET("/:id", controllers.ShowCliente)
			clientes.Use(middlewares.AuthAdminOrSupervisor()).PUT("/:id", controllers.UpdateCliente)
			clientes.Use(middlewares.AuthAdminOrSupervisor()).DELETE("/:id", controllers.DeleteCliente)
		}

		transportes := api.Group("transportes")
		{
			transportes.GET("/", controllers.ShowTransportes)
			transportes.Use(middlewares.AuthAdminOrSupervisor()).POST("/", controllers.CreateTransporte)
			transportes.Use(middlewares.AuthAdminOrSupervisor()).GET("/:id", controllers.ShowTransporte)
			transportes.Use(middlewares.AuthAdminOrSupervisor()).PUT("/:id", controllers.UpdateTransporte)
			transportes.Use(middlewares.AuthAdminOrSupervisor()).DELETE("/:id", controllers.DeleteTransporte)
		}

		cidades := api.Group("cidades")
		{
			cidades.GET("/", controllers.ShowCidades)
		}
	}

	return router
}
