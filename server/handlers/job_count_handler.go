package handlers

import (
	"encoding/json"
	"github.com/MattSilvaa/kleio/server/database"
	"github.com/MattSilvaa/kleio/server/models"
	"github.com/MattSilvaa/kleio/server/shared"
	"net/http"
	"time"
)

func FetchJobCount(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters from the request
	date := r.URL.Query().Get("date")
	location := r.URL.Query().Get("location")
	jobType := r.URL.Query().Get("jobType")

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusServiceUnavailable)
		return
	}

	// Validate job title and location using your validation functions
	if !models.IsValidJobTitle(jobType) || !models.IsValidLocation(location) {
		http.Error(w, "Invalid jobTitle or location", http.StatusBadRequest)
		return
	}

	// Fetch the job count from the database using the provided parameters
	jobCount, err := models.GetJobCountByParams(db, date, location, jobType)
	if err != nil {
		http.Error(w, "Failed to fetch job count", http.StatusInternalServerError)
		return
	}

	if jobCount == nil { // Assuming jobCount is of type *JobCount and you return nil if not found
		// Lock the critical section
		shared.ScrapeMutex.Lock()

		// Double-check if the data still doesn't exist in the DB (because another request might have inserted it)
		jobCount, _ = models.GetJobCountByParams(db, date, location, jobType)
		if jobCount == nil {
			count, err := models.GetJobCountFromScraper(jobType, location)
			if err != nil {
				http.Error(w, "Failed to scrape job counts", http.StatusInternalServerError)
				shared.ScrapeMutex.Unlock()
				return
			}

			err = models.InsertJobCount(db, location, jobType, count) // Modified InsertJobCount to accept these params
			if err != nil {
				http.Error(w, "Failed to insert job counts", http.StatusInternalServerError)
				shared.ScrapeMutex.Unlock()
				return
			}

			formattedDate, err := time.Parse("02-01-2006", date)
			if err != nil {
				http.Error(w, "Internal error", http.StatusInternalServerError)
				shared.ScrapeMutex.Unlock()
				return
			}

			jobCount = &models.JobCount{
				Date:  formattedDate,
				Count: count,
			}
		}

		// Unlock after the critical section
		shared.ScrapeMutex.Unlock()
	}

	// Convert the result to JSON and return the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobCount)
}
