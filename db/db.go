package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

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

	// Read environment variables for database configuration
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := "go-questions"
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	// Connection string setup
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
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

	_, err := Db.Exec(userTable)
	if err != nil {
		fmt.Println("Error creating users table:", err)
	}

	_, err = Db.Exec(questionsTable)
	if err != nil {
		fmt.Println("Error creating questions table:", err)
	} else {
		fmt.Println("Tables created or already exist!")
	}
}
