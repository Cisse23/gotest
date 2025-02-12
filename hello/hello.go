// Following tutorial https://go.dev/doc/tutorial/getting-started

// Package main is the entry point for a Go program.
package main

// Importing the fmt package for formatted I/O
import (
	"fmt"

	"log"

	"rsc.io/quote"

	"gotest/greetings"
)

// The main function is the starting point of the application.
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Declaring a string variable 'name' and initializing it with the value "World".
	var name string = "World"

	// Uncomment the following line to use short declaration syntax
	// for initializing multiple variables at once.
	// name, age, isAdult := "Matt", 23, true

	// Print a greeting message to the console.
	fmt.Println("Hello ", name)
	fmt.Println(quote.Go())

	// Request a greeting message.
	message, err := greetings.Hello("")

	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}

// Note:
// - To initialize the Go module, run: go mod init gotest/hello
// - To execute the program, run: go run hello.go
// - To clean up and add any missing module dependencies run: go mod tidy
