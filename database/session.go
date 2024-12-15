package database

import "fmt"

func CheckSession(sessionToken string) error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	var userID int
	err := db.QueryRow("SELECT user_id FROM session WHERE session = ?", sessionToken).Scan(&userID)
	if err != nil {
		return fmt.Errorf("invalid session token")
	}
	return nil
}
