package routes

import (
	"net/http"

	"office-user-management/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")

	router.HandleFunc("/api/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", handlers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	return router
}