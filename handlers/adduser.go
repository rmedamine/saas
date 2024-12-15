package handlers

import (
	"encoding/json"
	"net/http"
	"saas/database"
	"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Handler pour ajouter un utilisateur
func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User

		// Décoder les données JSON envoyées dans le corps de la requête
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Ajouter l'utilisateur à la base de données
		err = database.AddUser( user.Username, user.Password, user.Role)
		if err != nil {
			log.Printf("Failed to add user: %v", err)
			http.Error(w, "Failed to add user to database", http.StatusInternalServerError)
			return
		}

		// Répondre avec succès
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

