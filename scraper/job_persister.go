package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"

	"github.com/MattSilvaa/kleio/server/database"
	"github.com/MattSilvaa/kleio/server/models"
	"github.com/MattSilvaa/kleio/server/shared"
)

func JobPersister() error {
	db, err := database.ConnectToKleio()
	if err != nil {
		fmt.Printf("error connecting to database: %v", err)
		return err
	}
	defer db.Close()
	jobTitles := models.AllowedJobTitles
	locations := models.AllowedLocations

	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%02d-%02d-%d", day, month, year)

	g := new(errgroup.Group)
	for i := range locations {
		for j := range jobTitles {
			g.Go(func() error {
				location := locations[i]
				jobTitle := jobTitles[j]

				jobCount, err := models.GetJobCountByParams(db, date, location, jobTitle)
				if err != nil {
					fmt.Printf("failed to get job count from db: %v", err)
					return err
				}

				if jobCount == nil {
					shared.ScrapeMutex.Lock()
					defer shared.ScrapeMutex.Unlock()

					// Double-check if the data still doesn't exist in the DB (because another request might have inserted it)
					jobCount, _ = models.GetJobCountByParams(db, date, location, jobTitle)
					if jobCount == nil {
						count, err := models.GetJobCountFromScraper(jobTitle, location)
						if err != nil {
							fmt.Printf("failed to scrape job: %v", err)
							return err
						}

						err = models.InsertJobCount(db, location, jobTitle, count) // Modified InsertJobCount to accept these params
						if err != nil {
							fmt.Printf("Failed to insert job counts: %v", err)
							return err
						}

						formattedDate, err := time.Parse("02-01-2006", date)
						if err != nil {
							fmt.Printf("Failed to parse date: %v", err)
							return err
						}

						jobCount = &models.JobCount{
							Date:  &formattedDate,
							Count: &count,
						}
					}
				}
				return nil
			})
		}
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("received an error from a scrape job: %v", err)
		return err
	}

	return nil
}
