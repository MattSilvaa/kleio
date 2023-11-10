package models

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// JobCount structure to represent a job count entry in the database
type JobCount struct {
	ID       *int       `json:"id"`
	JobType  *string    `json:"job_type"`
	Location *string    `json:"location"`
	Count    *int       `json:"job_count"`
	Date     *time.Time `json:"date"`
}

// GetJobCountByParams fetches the job count based on the given parameters
func GetJobCountByParams(db *sql.DB, date, location, jobType string) (*JobCount, error) {
	query := `SELECT id, job_type, location, job_count, date FROM JobCount WHERE date = ? AND location = ? AND job_type = ?`
	row := db.QueryRow(query, date, location, jobType)

	jobCount := &JobCount{}
	err := row.Scan(jobCount.ID, jobCount.JobType, jobCount.Location, jobCount.Count, jobCount.Date)
	if err == sql.ErrNoRows {
		// No matching records found, but not an error condition
		return nil, nil
	} else if err != nil {
		// Other types of errors
		return nil, fmt.Errorf("failed to fetch job count: %v", err)
	}

	return jobCount, nil
}

// Runs python script and upload results to DB
func InsertJobCount(db *sql.DB, location string, jobType string, jobCount int) error {
	currentDate := time.Now().Format("2006-01-01")
	_, err := db.Exec("INSERT INTO JobCount (location, job_type, job_count, date) VALUES (?, ? , ?, ?)", location, jobType, jobCount, currentDate)
	return err
}

func GetJobCountFromScraper(jobType, location string) (int, error) {
	cmd := exec.Command("python3", "scraper/login_and_search.py", jobType, location)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to execute command: %s\nError: %v", cmd, err)
		return 0, nil
	}

	jobCountStr := string(output)
	// Regular expression to match a number with commas
	re := regexp.MustCompile(`\d{1,3}(,\d{3})*`)
	numberStr := re.FindString(jobCountStr)
	numberStr = strings.ReplaceAll(numberStr, ",", "")
	jobCount, err := strconv.Atoi(numberStr)
	if err != nil {
		log.Printf("Failed to convert job count string to integer: %s\nError: %v", jobCountStr, err)
		return 0, err
	}

	return jobCount, nil
}
