package database

import (
	"fmt"
	"log"
	"os/exec"
)

func GetTotalJobs() {
	cmd := exec.Command("python3", "scraper/login_and_search.py")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to execute command: %s\nError: %v", cmd, err)
	}
	fmt.Printf("Output: %s\n", output)
}
