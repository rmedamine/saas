package database

import "fmt"

func DeletSession(session string) error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	_, err := db.Exec("DELETE FROM session WHERE session = ?", session)
	if err != nil {
		return fmt.Errorf("failed to delete session: %v", err)
	}

	return nil
}
