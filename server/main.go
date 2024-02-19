package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MattSilvaa/kleio/server/database"
	"github.com/MattSilvaa/kleio/server/routes"
)

func main() {
	db, err := database.ConnectToSQL()
	if err != nil {
		log.Fatalf("Failed to connect to SQL: %v\n", err)
	}

	defer db.Close()

	err = database.CreateDB(db)
	if err != nil {
		log.Fatalf("Failed to create database: %v\n", err)
	}

	err = database.CreateTable(db)
	if err != nil {
		log.Fatalf("Failed to create table: %v\n", err)
	}

	router := routes.SetupRouter()

	// Create a server and set timeout values
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  25 * time.Second,
		WriteTimeout: 25 * time.Second,
	}

	// Start the server in a goroutine
	go func() {
		log.Println("Server listening on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v\n", err)
		}
	}()

	// Channel to listen for signals
	sigChan := make(chan os.Signal, 1)
	// Notify the sigChan on signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive the signal
	sigReceived := <-sigChan
	log.Printf("Received signal: %v. Shutting down...\n", sigReceived)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server stopped gracefully.")
}
