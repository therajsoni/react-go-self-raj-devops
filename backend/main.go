package main

import (
	"log"
	"net/http"
	"os"

	"office-user-management/database"
	"office-user-management/routes"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}

	database.ConnectMongoDB()

	router := routes.SetupRoutes()

	frontend_port := os.Getenv("FrontendPort")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			frontend_port,
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
		},
	}).Handler(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on " + port)

	err = http.ListenAndServe(":"+port, corsHandler)

	if err != nil {
		log.Fatal(err)
	}
}