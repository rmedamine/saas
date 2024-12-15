package middleware

import (
	"net/http"
	"saas/database"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Vérifiez le cookie de session
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Vérifiez si la session est valide
		err = database.CheckSession(cookie.Value)
		if err != nil {
			// Supprimez le cookie si la session est invalide
			http.SetCookie(w, &http.Cookie{
				Name:   "session",
				MaxAge: -1,
			})
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Passez au prochain gestionnaire
		next(w, r)
	}
}
