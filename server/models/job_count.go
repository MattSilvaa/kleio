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

// GetJobCountByParams fetches the job count based on the given parameters
func GetJobCountByParams(db *sql.DB, date, location, jobType string) (*JobCount, error) {
	query := `SELECT id, job_type, location, job_count, date FROM JobCount WHERE date = ? AND location = ? AND job_type = ?`
	row := db.QueryRow(query, date, location, jobType)

	var jobCount JobCount
	err := row.Scan(&jobCount.ID, &jobCount.JobType, &jobCount.Location, &jobCount.Count, &jobCount.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch job count: %v", err)
	}

	return &jobCount, nil
}

// Runs python script and upload results to DB
func InsertJobCount(db *sql.DB, location, jobCount string) error {
	currentDate := time.Now().Format("02-01-2006")
	_, err := db.Exec("INSERT INTO job_counts (date, count) VALUES (?, ?)", currentDate, count)
	return err
}

func GetJobCountFromScraper(jobType, location string) (int, error) {
	cmd := exec.Command("python3", "scraper/login_and_search.py", jobType, location)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to execute command: %s\nError: %v", cmd, err)
	}
	fmt.Printf("Output: %s\n", output)
}
