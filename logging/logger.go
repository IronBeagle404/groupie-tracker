package logging

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

// Open logfile for writing, create it if missing
// Define custom logger that writes to both console and logfile
func Init() {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file : %v", err)
	}

	mw := io.MultiWriter(os.Stdout, file)

	Logger = log.New(mw, "", log.LstdFlags)
}
