package main

import (
	"github.com/MattSilvaa/kleio/server/database"
	"log"
)

func main() {
	err := database.Create()
	if err != nil {
		log.Fatalf("Failed to create database:\t%v\n", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database:\t%v\n", err)
	}
	defer db.Close()
	database.GetTotalJobs()

	database.Delete()
}
