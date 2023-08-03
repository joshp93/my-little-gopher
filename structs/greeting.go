package structs

import "fmt"

type Greeting struct {
	Message string
	Err     error
}

func NewGreeting(randomFormat string, name string) *Greeting {
	return &Greeting{Message: fmt.Sprintf(randomFormat, name)}
}

func NewGreetingWithError(err error) *Greeting {
	return &Greeting{Err: err}
}
