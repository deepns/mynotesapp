package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestLoadNotes(t *testing.T) {
	// Create a temporary test file
	file, err := os.CreateTemp("", "testdata.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Create test data
	notes := []Note{
		{ID: 1, Description: "Note 1", Tags: []string{"tag1", "tag2"}, Contents: []string{"line1", "line2"}},
		{ID: 2, Description: "Note 2", Tags: []string{"tag1", "tag3"}, Contents: []string{"line1", "line2"}},
	}

	// Write test data to the test file
	err = json.NewEncoder(file).Encode(notes)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	// Load notes from the test file
	err = loadNotes(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Check if the loaded notes match the test data
	if !reflect.DeepEqual(notes, notes) {
		t.Errorf("Loaded notes do not match test data. Expected: %v, Got: %v", notes, notes)
	}
}

func TestFilterNotesByTags(t *testing.T) {
	notes := []Note{
		{ID: 1, Description: "Note 1", Tags: []string{"tag1", "tag2"}, Contents: []string{"line1", "line2"}},
		{ID: 2, Description: "Note 2", Tags: []string{"tag1", "tag3"}, Contents: []string{"line1", "line2"}},
		{ID: 3, Description: "Note 3", Tags: []string{"tag2", "tag3"}, Contents: []string{"line1", "line2"}},
	}

	// Filter notes by "tag1"
	filtered := filterNotesByTags(notes, []string{"tag1"})
	expected := []Note{
		{ID: 1, Description: "Note 1", Tags: []string{"tag1", "tag2"}, Contents: []string{"line1", "line2"}},
		{ID: 2, Description: "Note 2", Tags: []string{"tag1", "tag3"}, Contents: []string{"line1", "line2"}},
	}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Filtered notes do not match expected. Expected: %v, Got: %v", expected, filtered)
	}

	// Filter notes by "tag3"
	filtered = filterNotesByTags(notes, []string{"tag3"})
	expected = []Note{
		{ID: 2, Description: "Note 2", Tags: []string{"tag1", "tag3"}, Contents: []string{"line1", "line2"}},
		{ID: 3, Description: "Note 3", Tags: []string{"tag2", "tag3"}, Contents: []string{"line1", "line2"}},
	}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Filtered notes do not match expected. Expected: %v, Got: %v", expected, filtered)
	}

	// Filter notes by "tag4" (non-existing tag)
	filtered = filterNotesByTags(notes, []string{"tag4"})
	expected = []Note{}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Filtered notes do not match expected. Expected: %v, Got: %v", expected, filtered)
	}
}
