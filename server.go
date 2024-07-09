package main

import (
	"log"
	"net/http"
	"os"

	han "github.com/rodoben007/go-graphql-mongoDB/handler"
	"github.com/rodoben007/go-graphql-mongoDB/routes"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

var (
	GetDetails = han.GetDetails
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()
	routes.RestRoutes(router)
	routes.GraphRoutes(router)

	log.Printf("Server starting on port %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
