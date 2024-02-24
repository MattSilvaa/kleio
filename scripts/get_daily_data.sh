#!/bin/bash

# Path to your constants.go file
SCRAPER_MAIN="$PWD/scraper/main.go"
SCRAPER_JOB_SCRAPER="$PWD/scraper/job_persister.go"


# Activate virtual environment
source $PWD/../../pyenv/bin/activate

echo "Beginning jobs...."
go run $SCRAPER_MAIN $SCRAPER_JOB_SCRAPER
