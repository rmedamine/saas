	package database

func AddSession(session string, username string) error {
	user_id := 0
	err := db.QueryRow("SELECT user_id FROM users WHERE username = ?", username).Scan(&user_id)
	if err != nil {
		return err
	}
	_, _ = db.Exec("DELETE from session WHERE user_id = ?", user_id)

	_, err = db.Exec("INSERT INTO  session (session,user_id) VALUES (?,?)", session, user_id)
	return err
}
