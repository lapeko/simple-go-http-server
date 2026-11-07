//go:generate swag init -g cmd/api/main.go --dir ../../ --output ../../docs

package main

import (
	"counter/internal/api"
	"log"
	"os"
)

// @title           Counter API
// @version         1.0
// @description     Simple counter HTTP API.
// @host            localhost:3000
// @BasePath        /
func main() {
	port := os.Getenv("PORT")
	a := api.NewApi(port)
	a.Setup()
	log.Fatalln(a.Start())
}
