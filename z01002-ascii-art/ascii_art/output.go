package ascii_art

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Output(asciiText, outptFile string) {
	if filepath.Ext(outptFile) != ".txt" {
		fmt.Println(`Error: The output file should be text file`)
		os.Exit(0)
	} else if strings.Contains(outptFile, "/") {
		fmt.Println("Error: Wrong file format")
		os.Exit(0)
	}

	dir := "output"
    // Create the directory if it doesn't exist
    if err := os.MkdirAll(dir, 0755); err != nil {
        fmt.Println("failed to create directory: %w", err)
		os.Exit(1)
    }

    // Create the full file path
    filePath := filepath.Join(dir, outptFile)

    // Open the file, creating it if it doesn't exist, with read-write permissions
    file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        fmt.Println("failed to open file: %w", err)
		os.Exit(1)
    }
    defer file.Close()

    // Write the text to the file
    if _, err := file.WriteString(asciiText); err != nil {
        fmt.Println("failed to write to file: %w", err)
		os.Exit(1)
    }

    fmt.Println("Success: File successfully written in", filePath)

}

