package main

import (
	"github.com/MattSilvaa/kleio/server/database"
	"github.com/MattSilvaa/kleio/server/models"
	"log"
)

func main() {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		log.Fatalf("Failed to connect to db: %v\n", err)
	}

	err = database.CreateDB(db)
	if err != nil {
		log.Fatalf("Failed to create database: %v\n", err)
	}

	err = database.CreateTable(db)
	if err != nil {
		log.Fatalf("Failed to create table: %v\n", err)
	}

	models.UploadJobCount()
}
