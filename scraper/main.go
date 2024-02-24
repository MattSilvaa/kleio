package main

import (
	"fmt"
	"os"
)

func main() {
	err := JobPersister()
	if err != nil {
		fmt.Printf("Fatal error in job persister.\nExiting...")
		os.Exit(1)
	}
}
