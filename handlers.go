package myappnotes_function

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

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

func handleRoot(w http.ResponseWriter, r *http.Request) {
	randomNote := GetRandomNote()

	// Write the note as JSON response
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(randomNote)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNotes(w http.ResponseWriter, r *http.Request) {
	// Filter notes based on query parameters
	queryParams := r.URL.Query()
	tags := queryParams["tags"]
	limitStr := queryParams.Get("limit")

	filteredNotes := filterNotesByTags(tags)

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid 'limit' value", http.StatusBadRequest)
			return
		}

		if limit <= 0 {
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
