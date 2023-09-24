package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Create() error {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
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

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/kleio")
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to db")
	return db, nil
}

func Delete() error {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/kleio")
	if err != nil {
		return err
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS kleio")
	if err != nil {
		return err
	}

	fmt.Println("Successfully deleted db")
	return nil
}
