package server

import (
	"log"
	"tfweblog/routes"

	"github.com/gin-contrib/cors"
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

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.Use(cors.Default())
	log.Fatal(router.Run())
}
