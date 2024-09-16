package main

import (
	"log"
)

// handy to use these function as they will add newlines to the end
// as well as call exit or other functions as necessary

// may want to try package logrus at some point to get easy colors and warn etc.

// I put msg first to match log.Fatalf calls
func HandleError(message string, err error) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
