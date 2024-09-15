package main

import (
	"log"
)

// I put msg first to match log.Fatalf calls
func HandleError(message string, err error) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
