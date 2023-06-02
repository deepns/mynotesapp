package myappnotes_function

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	// Register the entry point function
	functions.HTTP("httpHandler", httpHandler)

	// Where to look for the notes.json file?
	sourceDir := "./"

	// Google Cloud Functions build places the source directory into
	// ./serverless_function_source_code in the packaged container
	// Set SOURCE_DIR=./serverless_function_source_code when running on GCF
	if envSourceDir := os.Getenv("SOURCE_DIR"); envSourceDir != "" {
		sourceDir = envSourceDir
	}

	// Hardcoding to notes.json for now
	// Fetch the data from github gist or use some db later
	// use path.Join() to join the sourceDir and notes.json
	notesFile := path.Join(sourceDir, "notes.json")

	err := loadNotes(notesFile)
	if err != nil {
		log.Fatal(err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %v", r.URL.Path)

	switch r.URL.Path {
	case "/":
		handleRoot(w, r)
	case "/notes":
		handleNotes(w, r)
	default:
		http.NotFound(w, r)
	}
}
