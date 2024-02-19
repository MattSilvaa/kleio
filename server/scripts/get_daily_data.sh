#!/bin/bash

# Path to your constants.go file
CONSTANTS_FILE="$PWD/server/models/constants.go"

JOB_TITLES=(
    "Artificial Intelligence (AI) Engineer",
    "Back-End Developer",
    "Business Analyst",
    "Cloud Solutions Architect",
    "Computer Hardware Engineer",
    "Cybersecurity Analyst",
    "Data Scientist",
    "Database Administrator (DBA)",
    "DevOps Engineer",
    "Front-End Developer",
    "Full-Stack Developer",
    "Game Developer",
    "IT Project Manager",
    "IT Support Specialist",
    "Machine Learning Engineer",
    "Mobile App Developer",
    "Network Administrator",
    "Network Engineer",
    "Product Designer",
    "Quality Assurance (QA) Engineer",
    "Software Engineer",
    "Systems Administrator",
    "Technical Writer",
    "UI/UX Designer",
    "Web Developer",
)

LOCATIONS=(
    "New York, NY",
    "Austin, TX",
    "Boston, MA",
    "Charlotte, NC",
    "Chicago, IL",
    "Columbus, OH",
    "Dallas, TX",
    "Denver, CO",
    "Detroit, MI",
    "El Paso, TX",
    "Fort Worth, TX",
    "Houston, TX",
    "Indianapolis, IN",
    "Jacksonville, FL",
    "Los Angeles, CA",
    "Nashville, TN",
    "Oklahoma City, OK",
    "Philadelphia, PA",
    "Phoenix, AZ",
    "San Antonio, TX",
    "San Diego, CA",
    "San Francisco, CA",
    "San Jose, CA",
    "Seattle, WA",
    "Washington, D.C.",
)

# Activate virtual environment
source $PWD/pyenv/bin/activate

# Iterate through job titles and locations
for JOB_TITLE in "$JOB_TITLES"; do
    for LOCATION in "$LOCATIONS"; do
        echo "Scraping for job title: $JOB_TITLE in location: $LOCATION"
        python $PWD/scraper/login_and_search.py "$JOB_TITLE" "$LOCATION"
    done
done

# Deactivate virtual environment
deactivate
