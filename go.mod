module gotest

go 1.22.2

//replace gotest/greetings => ../greetings
// The `replace` directive can be used to point a module or package path to a different location.
// In this project, it’s unnecessary because `greetings` is part of the same module (gotest).

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
