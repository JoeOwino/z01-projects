package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"ASCII-ART-WEB/asciiArt"
)

// Resizing
// file deletion
// error pages
func main() {
	if len(os.Args) != 1 {
		fmt.Printf("usage: go run main.go" + "\n")
		return
	}

	// checksum results, to ascertain non of the banner files is tampered with
	results := asciiArt.VerifyBanners()

	for filePath, isValid := range results {
		if isValid {
			//	fmt.Printf("Checksum for %s is valid.\n", filePath)
		} else {
			fmt.Printf("Checksum for %s is invalid.\n", filePath)
			return
		}
	}
	http.HandleFunc("/", pathHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("  Server listening on http://localhost:8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path

	if filePath == "/" {
		http.ServeFile(w, r, "./static")
	} else if strings.HasPrefix(filePath, "/static/") {
		http.ServeFile(w, r, "."+filePath)
	} else {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/err404.html")
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "static/err405.html")
		return
	}

	var request struct {
		Banner string `json:"banner"`
		Input  string `json:"input"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "400 - Invalid request body", http.StatusBadRequest)
		return
	}

	fileName := asciiArt.BannerFile(request.Banner)

	bannerMap, err := asciiArt.LoadBannerMap(fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "500 - Failed to load banner file", http.StatusInternalServerError)
		return
	}

	response, httpErr := generateASCIIArt(w, request.Input, bannerMap)
	if httpErr != nil {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}

// generateASCIIArt generates the ASCII art from input and bannerMap.
// Logs an error using http.Error if input contains characters outside ASCII limits.
func generateASCIIArt(w http.ResponseWriter, input string, bannerMap map[int][]string) (string, error) {
	var str strings.Builder
	lines := make([]string, 8)
	input = strings.ReplaceAll(input, "\r", "\n")
	arr := strings.Split(input, "\n")
	for _, line := range arr {
		for _, char := range line {

			banner, exists := bannerMap[int(char)]
			if !exists {
				w.WriteHeader(http.StatusBadRequest)
				http.Error(w, fmt.Sprintf("400 - Character '%c' not found in banner map", char), http.StatusBadRequest)
				return "", fmt.Errorf("character '%c' not found in banner map", char)
			}
			for i := 0; i < 8; i++ {
				lines[i] += banner[i]
			}
		}
		str.WriteString(strings.Join(lines, "\n"))
		str.WriteString("\n")
		lines = make([]string, 8)
	}
	return str.String(), nil
}
