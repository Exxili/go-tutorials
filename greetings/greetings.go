package greetings

import "fmt"

// Returns a greeting for name passed in arg
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
