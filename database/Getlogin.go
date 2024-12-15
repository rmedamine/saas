package database

import "database/sql"

func GetLogin(username, password string) (bool, error) {
	if db == nil {
		return false, sql.ErrConnDone
	}

	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // Aucun utilisateur trouvé
		}
		return false, err // Autre erreur
	}

	// Comparer le mot de passe directement (ajoutez bcrypt si nécessaire)
	if password == storedPassword {
		return true, nil
	}

	return false, nil
}
