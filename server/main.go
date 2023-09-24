package main

import (
	"log"
)

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatalf("Error connecting to the database:\t%v\n", err)
	}
	defer db.Close()
}
