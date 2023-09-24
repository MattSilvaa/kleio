package db

import "database/sql"

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/database-name")

	if err != nil {
		return nil, err
	}

	return db, nil
}
