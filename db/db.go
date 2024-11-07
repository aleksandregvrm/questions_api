package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDatabase() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error occurred in .env file; please check.")
	}

	dbUrl := os.Getenv("DATABASE_URL")

	db, errSql := sql.Open("postgres", dbUrl)
	if errSql != nil {
		fmt.Println("Error connecting to the database:", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}

	// Create tables if they do not exist
	createTables()
}

func createTables() {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		password VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	questionsTable := `
	CREATE TABLE IF NOT EXISTS questions (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		question_text TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	userCreds := `
	CREATE TABLE IF NOT EXISTS credentials (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		country VARCHAR(100) NOT NULL,
		city VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS user_registrations (
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	credentials_id INTEGER,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (credentials_id) REFERENCES credentials(id)
	)
	`

	// Users Table
	_, err := Db.Exec(userTable)
	if err != nil {
		fmt.Println("Error creating users table:", err)
	}

	// Credentials Table
	_, err = Db.Exec(userCreds)
	if err != nil {
		fmt.Println("Error creating credentials table:", err)
	}

	// Questions Table
	_, err = Db.Exec(questionsTable)
	if err != nil {
		fmt.Println("Error creating questions table:", err)
	} else {
		fmt.Println("Tables created or already exist!")
	}

	// user_cred registration
	_, err = Db.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}
}
