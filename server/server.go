package server

import (
	"log"
	"tfweblog/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		server: gin.Default(),
	}
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://tfweblog-frontend.herokuapp.com")
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

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.Use(CORSMiddleware())
	log.Fatal(router.Run())
}
