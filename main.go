package main

import (
	"apiGo/database"
	"apiGo/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
