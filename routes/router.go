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

		dashboard := api.Group("dashboard", middlewares.Auth())
		{
			dashboard.GET("/", controllers.Dashboard)
		}

		usuarios := api.Group("usuarios", middlewares.AuthAdmin())
		{
			usuarios.GET("/", controllers.ShowUsuarios)
			usuarios.POST("/", controllers.CreateUsuario)
			usuarios.GET("/:id", controllers.ShowUsuario)
			usuarios.PUT("/:id", controllers.UpdateUsuario)
			usuarios.DELETE("/:id", controllers.DeleteUsuario)
		}
	}

	return router
}
