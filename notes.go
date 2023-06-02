package myappnotes_function

import (
	"math/rand"
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

func init() {
	// Seed the random number generator with the current timestamp
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRandomNote() Note {
	return notes[rng.Intn(len(notes))]
}

func filterNotesByTags(filterTags []string) []Note {
	if len(filterTags) <= 0 {
		return notes
	}

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
