package main

import (
	"log"
	"net/http"

	"github.com/MagnusHLund/VisitorCounter/internal/routes"
	"github.com/MagnusHLund/VisitorCounter/internal/wire"
)

func main() {
	app, err := wire.CreateApp()
	if err != nil {
		log.Fatal("Failed to create app: ", err)
	}

	router := routes.SetupRoutes(app)

	startServer(router)
}

func startServer(router http.Handler) {
	var port string = "8080" // TODO: use environment variables

	log.Println("Starting server on port", port+"...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
