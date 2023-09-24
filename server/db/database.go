package database

import (
	"database/sql"
	"fmt"
)

func create() error {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS kleio")
	if err != nil {
		return err
	}

	fmt.Println("Database 'kleio' created (if it didn't exist)")
	return nil
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/kleio")
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
