package main

import (
	"net/http"

	"./controller"
)

func main() {
	// Using custom Multiplexer to handle the routes
	mux := http.NewServeMux()
	// Configuring the routes in main controller
	controller.ConfigureRoutes(mux)
	http.ListenAndServe(":3000", mux)
}
