package main

import (
	"os"
	"Golangweb/routes"

)

func main() {
	var appPort = os.Getenv("appPort")

	if appPort == "" {
		appPort = "8080"
	}

	routes.SetupServer(appPort)
}