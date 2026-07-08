package main

import (
	"bufio"
	"fmt"
	"go-reloaded/functions"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {

		fmt.Println("Usage: go run . <input file> <output file>")
		return
	}

	inputFile := os.Args[1]

	outputFile := os.Args[2]

	content, err := readFile(inputFile)

	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	processed := functions.ProcessText(content)

	err = writeFile(outputFile, processed)

	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return // Exit the program
	}

	// Print success message after processing and saving the file
	fmt.Printf("Successfully processed %s to %s\n", inputFile, outputFile)
}

// readFile reads the content of a file and returns it as a string
func readFile(filename string) (string, error) {
	// Open the file with the given filename
	file, err := os.Open(filename)

	// Return an error if the file cannot be opened
	if err != nil {
		return "", err
	}

	// Ensure the file is closed after the function finishes
	defer file.Close()

	// Create a string builder to efficiently build the file content
	var content strings.Builder

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through each line in the file
	for scanner.Scan() {
		// Append the current line to the string builder
		content.WriteString(scanner.Text())

		// If the line is not empty, append a newline character
		if scanner.Text() != "" {
			content.WriteString("\n")
		}
	}

	// Return the final content after trimming trailing newline characters
	return content.String(), scanner.Err()
}

// writeFile writes the provided content string to the specified file
func writeFile(filename, content string) error {
	// Write the content to the file with permission 0644
	return os.WriteFile(filename, []byte(content), 0644)
}
