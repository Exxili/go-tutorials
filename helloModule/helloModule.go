package main

import (
	"fmt"

	"./greetings"
)

func main() {
	// Get a greeting message
	message := greetings.Hello("Andy")
	fmt.Println(message)
}
