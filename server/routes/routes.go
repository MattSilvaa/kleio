package routes

import (
	"net/http"

	"github.com/MattSilvaa/kleio/server/handlers"
	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // or specify your frontend's URL instead of '*'
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(enableCORS)

	// Define a route for fetching the job count based on date, location, and job type
	router.HandleFunc("/api/job-count", handlers.FetchJobCount).Methods(http.MethodGet)

	return router
}
