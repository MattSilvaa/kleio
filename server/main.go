package main

import (
	database "github.com/MattSilvaa/kleio/server/database"
	"log"
)

func main() {
	db, err := database.connect()
	if err != nil {
		log.Fatalf("Error connecting to the database:\t%v\n", err)
	}
	defer db.Close()
}
