// Application main package.
package main

import (
	"CartService/pkg/controllers"
	"CartService/pkg/filters"
	"CartService/pkg/routes"
	"CartService/pkg/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
)

// controllers
var cartController controllers.CartController

// services
var userService services.HttpUserService
var cartService services.InMemoryCartService

// middleware
var securityFilter filters.SecurityFilter

// routes
var cartRoutes routes.CartRoutes

// Application entry point.
// Sets up dependcy injection, routing and starts the HTTP server.
func main() {

	setupDependencyGraph()

	router := setupRoutes()

	startMiddleware(router)

	startServer(router)
}

// Populates the dependeny graph for depdency injection.
func setupDependencyGraph() {

	var graph inject.Graph

	err := graph.Provide(
		// controllers
		&inject.Object{Value: &cartController},

		// services
		&inject.Object{Value: &userService},
		&inject.Object{Value: &cartService},

		// filters
		&inject.Object{Value: &securityFilter},

		// rotues
		&inject.Object{Value: &cartRoutes},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := graph.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Setup the web servce routing.
func setupRoutes() *mux.Router {

	router := mux.NewRouter()

	// cart
	cartRoutes.Create(router)

	return router
}

// Adds the routing middleware.
func startMiddleware(router *mux.Router) {
	router.Use(securityFilter.Execute)
}

// Adds the router and starts the HTTP server on a default port of 8080.
func startServer(router *mux.Router) {

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 25 * time.Second,
		ReadTimeout:  25 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
