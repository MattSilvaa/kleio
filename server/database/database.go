package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	env := os.Getenv("KLEIO_ENV")

	var dsn string
	if env == "test" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/kleio",
			os.Getenv("LOCAL_KLEIO_USER"),
			os.Getenv("LOCAL_KLEIO_PW"),
			os.Getenv("LOCAL_KLEIO_HOST"),
		)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/kleio",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
		)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Printf("Successfully connected.\n")
	return db, nil
}

func CreateDB(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS kleio")
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	fmt.Println("Database 'kleio' created (if it didn't exist)")
	return nil
}

func CreateTable(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS kleio.JobCount (
		id INT AUTO_INCREMENT PRIMARY KEY,
		job_type VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		job_count INT NOT NULL,
		date DATE NOT NULL
	);`

	if _, err := db.Exec(createTableSQL); err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	fmt.Println("Table created successfully (if it didn't exist)")
	return nil
}

func DeleteDB(db *sql.DB) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS kleio")
	if err != nil {
		return fmt.Errorf("error deleting database: %v", err)
	}

	fmt.Println("Successfully deleted db")
	return nil
}
