package server

import (
	"io"
	"log"
	"os"
)

var (
	Debug *log.Logger
)

func InitLogger(debug bool) {
	// Discard policy by default
	Debug = log.New(io.Discard, "", 0)
	// Enable it only if debug flag was passed
	if debug {
		Debug = log.New(os.Stderr, "DEBUG: ", log.Ltime|log.Lshortfile)
	}
}
