package handlers

import (
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	// Routes statiques pour les fichiers CSS, JS, etc.
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Routes publiques
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/", GetHome)
	mux.HandleFunc("/add-user", AddUser)


	// Routes protégées avec middleware

	return mux
}
