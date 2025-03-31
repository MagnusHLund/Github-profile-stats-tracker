package main

import (
	"log"
	"net/http"

	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/routes"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/wire"
)

func main() {
	app, err := wire.CreateApp()
	if err != nil {
		log.Fatal("Failed to create app: ", err)
	}

	port := app.Config.Config.ServerConfig.ServerPort
	router := routes.SetupRoutes(app)

	startServer(router, port)
}

func startServer(router http.Handler, port string) {
	log.Println("Starting server on port", port+"...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
