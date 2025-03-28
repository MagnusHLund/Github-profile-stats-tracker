package main

import (
	"log"
	"net/http"

	"github.com/MagnusHLund/VisitorCounter/internal/handlers"
	"github.com/MagnusHLund/VisitorCounter/internal/routes"
	"github.com/MagnusHLund/VisitorCounter/internal/wire"
)

func main() {
	app, err := wire.CreateApp()
	if err != nil {
		log.Fatal("Failed to create app: ", err)
	}

	router := routes.SetupRoutes(&app.Handlers)
	startServer(router)
}

func startServer(router http.Handler) {
	const port = "8080"

	log.Println("Starting server on port", port+"...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
