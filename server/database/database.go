package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}
	return nil
}

// getDSN constructs and returns the Data Source Name (DSN) string based on the current environment.
func getDSN(useKleio bool) (string, error) {
	if err := loadEnv(); err != nil {
		return "", err
	}

	env := os.Getenv("ENV")
	var schema string
	if useKleio {
		schema = "kleio"
	}

	var dsn string
	switch env {
	case "test":
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s",
			os.Getenv("LOCAL_KLEIO_USER"),
			os.Getenv("LOCAL_KLEIO_PW"),
			os.Getenv("LOCAL_KLEIO_HOST"),
			schema,
		)
	default:
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			schema,
		)
	}

	return dsn, nil
}

func ConnectToSQL() (*sql.DB, error) {
	dsn, err := getDSN(false)
	if err != nil {
		log.Printf("Error getting dsn: %v", err)
		return nil, err
	}

	// Open a new database connection.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}
	// Test the database connection.
	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database.")
	return db, nil
}

func ConnectToKleio() (*sql.DB, error) {
	dsn, err := getDSN(true)

	if err != nil {
		log.Printf("Error getting dsn: %v", err)
		return nil, err
	}
	// Open a new database connection.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}
	// Test the database connection.
	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database.")
	return db, nil
}

// CreateDB creates the database if it does not exist.
func CreateDB(db *sql.DB) error {
	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS kleio"); err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	log.Println("Database 'kleio' created (if it didn't exist)")
	return nil
}

// CreateTable creates a new table in the database if it does not exist.
func CreateTable(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS JobCount (
		id INT AUTO_INCREMENT PRIMARY KEY,
		job_type VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		job_count INT NOT NULL,
		date DATE NOT NULL
	);`

	if _, err := db.Exec("USE kleio;"); err != nil {
		return fmt.Errorf("error using kleio: %v", err)
	}

	if _, err := db.Exec(createTableSQL); err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	log.Println("Table 'JobCount' created successfully (if it didn't exist)")
	return nil
}

// DeleteDB drops the database.
func DeleteDB(db *sql.DB) error {
	if _, err := db.Exec("DROP DATABASE IF EXISTS kleio"); err != nil {
		return fmt.Errorf("error deleting database: %v", err)
	}

	log.Println("Successfully deleted database 'kleio'")
	return nil
}
