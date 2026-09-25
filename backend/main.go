package main

import (
	"log"
	"net/http"
	"os"

	"office-user-management/database"
	"office-user-management/routes"

	"github.com/rs/cors"
)

func main() {

	database.ConnectMongoDB()

	router := routes.SetupRoutes()

	frontendPort := os.Getenv("FrontendPort")

	if frontendPort == "" {
		frontendPort = "http://localhost"
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			frontendPort,
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	}).Handler(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port " + port)

	err := http.ListenAndServe(":"+port, corsHandler)

	if err != nil {
		log.Fatal(err)
	}
}