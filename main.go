package main // Defines the main package. Every executable Go program must have package main.

import (
	"bufio"                 // Provides buffered I/O utilities (used for reading files efficiently)
	"fmt"                   // Provides formatted I/O functions like Println and Printf
	"go-reloaded/functions" // Imports a custom package that contains the ProcessText function
	"os"                    // Provides functions for interacting with the operating system (files, arguments, etc.)
	"strings"               // Provides string manipulation utilities
)

func main() { // Entry point of the program
	// Check if the number of command-line arguments is not equal to 3
	// (program name + input file + output file)
	if len(os.Args) != 3 {
		// Print usage instruction if arguments are incorrect
		fmt.Println("Usage: go run . <input file> <output file>")
		return // Exit the program
	}

	// Store the first command-line argument as the input file name
	inputFile := os.Args[1]

	// Store the second command-line argument as the output file name
	outputFile := os.Args[2]

	// Read the content of the input file
	content, err := readFile(inputFile)

	// Check if an error occurred while reading the file
	if err != nil {
		// Print the error message
		fmt.Printf("Error reading input file: %v\n", err)
		return // Exit the program
	}

	// Process the text content using a custom function from the functions package
	processed := functions.ProcessText(content)

	// Write the processed content into the output file
	err = writeFile(outputFile, processed)

	// Check if an error occurred while writing the file
	if err != nil {
		// Print the error message
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
