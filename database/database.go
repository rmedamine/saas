package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB // Variable globale pour la connexion à la base de données

// InitializeDB initialise la connexion à la base de données
func InitializeDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
		return err
	}

	// Vérifiez la connexion
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	CreateDb()

	log.Printf("Database connected successfully: %s", dataSourceName)
	return nil
}

// CreateDb crée les tables nécessaires
func CreateDb() {
	// Assurez-vous que la base de données est connectée
	if db == nil {
		log.Fatal("Database connection is not initialized.")
	}

	log.Println("Creating tables...")

	// Création de la table 'users'
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(255) UNIQUE NOT NULL,
		role VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	`)
	if err != nil {
		log.Fatalf("Failed to create 'users' table: %v", err)
	} else {
		log.Println("'users' table created successfully.")
	}

	// Création de la table 'session'
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS session (
		session TEXT,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		log.Fatalf("Failed to create 'session' table: %v", err)
	} else {
		log.Println("'session' table created successfully.")
	}

	// Insérer des données de test avec gestion des conflits
	log.Println("Inserting test data into 'users' table...")
	_, err = db.Exec(`
		INSERT OR IGNORE INTO users (username, role, password) VALUES 
			('aaa', 'admin', '123');
	`)
	if err != nil {
		log.Printf("Failed to insert test data: %v", err)
	} else {
		log.Println("Test data inserted successfully.")
	}
}

// CloseDB ferme la connexion à la base de données
func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Printf("Failed to close database: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	}
}
