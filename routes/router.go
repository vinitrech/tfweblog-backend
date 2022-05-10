package routes

import (
	"tfweblog/controllers"
	"tfweblog/server/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://tfweblog-frontend.herokuapp.com"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://tfweblog-frontend.herokuapp.com"
		},
		MaxAge: 24 * time.Hour,
	}))

	api := router.Group("api/v1")
	{

		api.POST("login", controllers.Login)
		api.POST("cadastro", controllers.Cadastro)
		api.POST("auth/google", controllers.GoogleLogin)
		api.GET("/getData", controllers.GetData)

		api.Use(middlewares.Auth())

		api.GET("/dashboard", cors.New(cors.Config{
			AllowOrigins:     []string{"https://tfweblog-frontend.herokuapp.com"},
			AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "DELETE", "GET", "POST"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://tfweblog-frontend.herokuapp.com"
			},
			MaxAge: 24 * time.Hour,
		}), controllers.Dashboard)

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
			transportes.GET("/", controllers.ShowTransportes)
			transportes.POST("/", middlewares.AuthAdminOrSupervisor(), controllers.CreateTransporte)
			transportes.GET("/:id", middlewares.AuthAdminOrSupervisor(), controllers.ShowTransporte)
			transportes.PUT("/:id", middlewares.AuthAdminOrSupervisor(), controllers.UpdateTransporte)
			transportes.DELETE("/:id", middlewares.AuthAdminOrSupervisor(), controllers.DeleteTransporte)
			transportes.GET("/enviar-inicio-supervisor/:id", middlewares.AuthAdminOrDriver(), controllers.EnviarInicioSupervisor)
			transportes.GET("/aprovar-inicio/:id", middlewares.AuthAdminOrSupervisor(), controllers.AprovarInicio)
			transportes.GET("/enviar-finalizacao-supervisor/:id", middlewares.AuthAdminOrDriver(), controllers.EnviarFinalizacaoSupervisor)
			transportes.GET("/finalizar/:id", middlewares.AuthAdminOrSupervisor(), controllers.Finalizar)
		}

		cidades := api.Group("cidades")
		{
			cidades.GET("/", controllers.ShowCidades)
		}
	}

	return router
}
