package main

import (
	"log"

	_ "github.com/pealan/golang-rest-api/docs"
)

// @title           Golang Rest API
// @version         1.0
// @description     A device management service REST API in Go using Gin framework.

// @contact.name   Pedro Ramos
// @contact.email  pealan97@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {
	server, err := InitializeAPI()
	if err != nil {
		log.Fatalf("Could not initiate server %+v", err)
	}

	server.Start()
}
