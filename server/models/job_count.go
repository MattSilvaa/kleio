package models

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"time"
)

// JobCount structure to represent a job count entry in the database
type JobCount struct {
	ID       int       `json:"id"`
	JobType  string    `json:"job_type"`
	Location string    `json:"location"`
	Count    int       `json:"job_count"`
	Date     time.Time `json:"date"`
}

// DB is the database connection instance (you can initialize this in your main function or another setup function)
var DB *sql.DB

// GetJobCountByParams fetches the job count based on the given parameters
func GetJobCountByParams(date, location, jobType string) (*JobCount, error) {
	query := `SELECT id, job_type, location, job_count, date FROM JobCount WHERE date = ? AND location = ? AND job_type = ?`
	row := DB.QueryRow(query, date, location, jobType)

	var jobCount JobCount
	err := row.Scan(&jobCount.ID, &jobCount.JobType, &jobCount.Location, &jobCount.Count, &jobCount.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch job count: %v", err)
	}

	return &jobCount, nil
}

// Runs python script and upload results to DB
func UploadJobCount() {
	cmd := exec.Command("python3", "scraper/login_and_search.py")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to execute command: %s\nError: %v", cmd, err)
	}
	fmt.Printf("Output: %s\n", output)
}
