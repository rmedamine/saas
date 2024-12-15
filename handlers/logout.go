package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"saas/database"
)

// Structure pour gérer les données de déconnexion
type Logout struct {
	Session string `json:"session"`
}

// Logout_S gère la déconnexion de l'utilisateur
func Logout_S(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Essayer de récupérer la session depuis le cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		// Si aucun cookie, essayer de lire les données JSON
		var logout Logout
		err := json.NewDecoder(r.Body).Decode(&logout)
		if err != nil || logout.Session == "" {
			log.Println("No valid session provided")
			http.Error(w, "Invalid session data", http.StatusBadRequest)
			return
		}

		// Utiliser la session depuis les données JSON
		err = database.DeletSession(logout.Session)
		if err != nil {
			log.Printf("Failed to delete session: %v", err)
			http.Error(w, "Failed to logout", http.StatusInternalServerError)
			return
		}

		// Réponse de succès
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logout successful"))
		return
	}

	// Supprimer la session basée sur le cookie
	err = database.DeletSession(cookie.Value)
	if err != nil {
		log.Printf("Failed to delete session: %v", err)
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	// Supprimer le cookie en envoyant un cookie expiré
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	// Réponse de succès
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
