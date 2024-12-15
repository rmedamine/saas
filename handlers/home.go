package handlers

import (
	"net/http"
)

// GetHome affiche la page d'accueil si l'utilisateur est authentifié
func GetHome(w http.ResponseWriter, r *http.Request) {
	// Récupérez le cookie de session
	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" {
		// Si aucun cookie de session n'existe ou est invalide, redirigez vers la page de connexion
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Si la session est valide, affichez la page d'accueil
	// Vous pouvez également afficher le nom de l'utilisateur ou d'autres données si nécessaire
	http.ServeFile(w, r, "./assets/templates/home.html")
}
