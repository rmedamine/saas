package main

import (
	"log"
	"net/http"
	"saas/database"
	"saas/handlers"
)


func main() {
	// Initialisation de la base de données
	err := database.InitializeDB("./test.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CloseDB()

	// Configuration du serveur HTTP
	srv := http.Server{
		Addr:    ":8080",
		Handler: handlers.Routes(),
	}

	log.Println("Server running on http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
