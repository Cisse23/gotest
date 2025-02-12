package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// se the go mod edit command to edit the gotest/hello module to redirect
// Go tools from its module path (where the module isn't) to the local directory (where it is).
// run: $ go mod edit -replace gotest/greetings=../greetings
