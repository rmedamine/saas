package database

import (
	"database/sql"
	"fmt"
	"log"
)

func AddUser(username, password, role string) error {
	// Vérifie si le nom d'utilisateur existe déjà
	var existingUsername string
	err := db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&existingUsername)
	if err != sql.ErrNoRows {
		if err == nil {
			// L'utilisateur existe déjà
			return fmt.Errorf("username %s already taken", username)
		}
		log.Printf("Error checking username: %v", err)
		return err
	}

	// Ajouter l'utilisateur si le nom d'utilisateur est unique
	_, err = db.Exec(`
		INSERT INTO users (username, password, role) 
		VALUES (?, ?, ?)`,
		username, password, role)
	if err != nil {
		log.Printf("Error adding user: %v", err)
		return err
	}

	log.Println("User added successfully.")
	return nil
}
