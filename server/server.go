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

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Fatal(router.Run())
}
