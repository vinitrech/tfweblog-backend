package main

import (
	"tfweblog/database"
	"tfweblog/server"
)

func main() {
	database.StartDB()
	
	server := server.NewServer()

	server.Run()
}
