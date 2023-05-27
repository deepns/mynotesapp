package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Note struct {
	ID          int      `json:"id"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Contents    []string `json:"contents"`
}

var notes []Note
var rng *rand.Rand

func main() {
	// Check the number of command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <testfile> <port>")
		return
	}

	// Get the command line arguments
	testFile := os.Args[1]
	port := os.Args[2]

	// Load notes from the JSON file
	err := loadNotes(testFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the random number generator
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define the HTTP routes and their respective handlers
	http.HandleFunc("/", getRandomNoteHandler)
	http.HandleFunc("/notes", getAllNotesHandler)

	// Start the web server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting web server on http://localhost%s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loadNotes(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&notes)
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func getRandomNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Generate a random index
	randomIndex := rng.Intn(len(notes))

	// Retrieve the random note
	randomNote := notes[randomIndex]

	// Write the note as JSON response
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(randomNote)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func getAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Filter notes based on query parameters
	queryParams := r.URL.Query()
	tags := queryParams["tags"]
	limitStr := queryParams.Get("limit")

	filteredNotes := notes
	if len(tags) > 0 {
		filteredNotes = filterNotesByTags(filteredNotes, tags)
	}
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid 'limit' value", http.StatusBadRequest)
			return
		}
		filteredNotes = limitNotes(filteredNotes, limit)
	}

	// Write the filtered notes as JSON response
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(filteredNotes)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func filterNotesByTags(notes []Note, filterTags []string) []Note {
	filtered := make([]Note, 0)
	for _, note := range notes {
		if containsAnyTag(note.Tags, filterTags) {
			filtered = append(filtered, note)
		}
	}
	return filtered
}

func containsAnyTag(noteTags []string, filterTags []string) bool {
	for _, filterTag := range filterTags {
		for _, noteTag := range noteTags {
			if strings.EqualFold(noteTag, filterTag) {
				return true
			}
		}
	}
	return false
}

func limitNotes(notes []Note, limit int) []Note {
	if limit >= len(notes) {
		return notes
	}
	return notes[:limit]
}
