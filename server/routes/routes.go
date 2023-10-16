package routes

import (
	"github.com/MattSilvaa/kleio/server/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Define a route for fetching the job count based on date, location, and job type
	router.HandleFunc("/api/jobcount", handlers.FetchJobCount).Methods(http.MethodGet)

	return router
}
