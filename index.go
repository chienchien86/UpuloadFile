package main

import (
	"os"
	"UpuloadFile/routes"

)

func main() {
	var appPort = os.Getenv("appPort")

	if appPort == "" {
		appPort = "8080"
	}

	routes.SetupServer(appPort)
}