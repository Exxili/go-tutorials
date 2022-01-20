package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	// Get a greeting message
	message := greetings.Hello("Andy")
	fmt.Println(message)
}
