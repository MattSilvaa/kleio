package handlers

import (
	"encoding/json"
	"github.com/MattSilvaa/kleio/server/models"
	"net/http"
)

func FetchJobCount(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters from the request
	date := r.URL.Query().Get("date")
	location := r.URL.Query().Get("location")
	jobType := r.URL.Query().Get("jobType")

	// Fetch the job count from the database using the provided parameters
	jobCount, err := models.GetJobCountByParams(date, location, jobType)
	if err != nil {
		http.Error(w, "Failed to fetch job count", http.StatusInternalServerError)
		return
	}

	// Convert the result to JSON and return the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobCount)
}
