package main

import (
	"VisitorCounter/internal/routes"
	"VisitorCounter/internal/wire"
	"log"
	"net/http"
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
	const port = "8080"

	log.Println("Starting server on port", port+"...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
