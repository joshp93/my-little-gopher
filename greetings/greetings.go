package greetings

import (
	"errors"
	"math/rand"
	"my-little-gopher/structs"
)

func Hello(name string) *structs.Greeting {
	if name == "" {
		return structs.NewGreetingWithError(errors.New("empty name"))
	}
	greeting := structs.NewGreeting(randomFormat(), name)
	return greeting
}

func Hellos(names []string, hellosChan chan<- string, errorChan chan<- error) {
	for _, name := range names {
		greeting := Hello(name)
		if greeting.Err != nil {
			errorChan <- greeting.Err
			break
		}
		hellosChan <- greeting.Message
	}
	close(errorChan)
	close(hellosChan)
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Go away %v, you are clearly a rustacean and are therefore not welcome here",
	}

	return formats[rand.Intn(len(formats))]
}
