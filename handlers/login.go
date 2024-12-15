package handlers

import (
	"log"
	"net/http"
	"saas/bcryptp"
	"saas/database"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login handler reached")

	if r.Method == http.MethodGet {
		// Servir la page de connexion
		http.ServeFile(w, r, "./assets/templates/login.html")
	} else if r.Method == http.MethodPost {
		log.Println("Processing login POST request")

		// Récupérez les données du formulaire
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Vérifiez les identifiants
		isAuthenticated, err := database.GetLogin(username, password)
		if err != nil {
			log.Printf("Error during login: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if isAuthenticated {
			// Générer un token de session unique
			sessionToken, err := bcryptp.CreateSession()
			if err != nil {
				log.Printf("Failed to create session token: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Ajouter la session à la base de données
			err = database.AddSession(sessionToken.String(), username)
			if err != nil {
				log.Printf("Failed to add session to database: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Créez un cookie de session
			http.SetCookie(w, &http.Cookie{
				Name:  "session",
				Value: sessionToken.String(),
				Path:  "/",
			})

			// Servir directement la page d'accueil après la connexion
			// http.ServeFile(w, r, "./assets/templates/home.html")
			http.Redirect(w, r, "/home", http.StatusSeeOther)

		} else {
			// Identifiants invalides
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	}
}
