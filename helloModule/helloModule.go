package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {
	// Set a prefix for the logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message
	message, err := greetings.Hello("")

	// If there is an error, print to console and exit
	if err != nil {
		log.Fatal(err)
	}

	// if no Err, print returned message
	fmt.Println(message)
}
