package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string, hellosChan chan<- string, errorChan chan<- error) {
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			errorChan <- err
			break
		}
		hellosChan <- message
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
