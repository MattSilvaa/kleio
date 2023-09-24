package database

import "database/sql"

func create() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/database-name")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/database-name")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func delete() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/database-name")

	if err != nil {
		return nil, err
	}

	return db, nil
}
