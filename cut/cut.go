package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Please provide a valid column number and at least one argument.")
		return
	}

	// Parse arguments
	desiredField, fileToOpen, delimiter, err := parseArguments(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Process the file
	if err := processFile(fileToOpen, desiredField, delimiter); err != nil {
		fmt.Println("Error processing file:", err)
	}
}
