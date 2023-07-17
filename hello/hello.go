package hello

import (
	"fmt"
	"log"

	"my-little-gopher/greetings"
)

func SayHello() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	hellosChan := make(chan string, len(names))
	errorChan := make(chan error, 1)

	go greetings.Hellos(names, hellosChan, errorChan)
	err := <-errorChan
	if err != nil {
		log.Fatal(err)
	}

	for message := range hellosChan {
		fmt.Println(message)
	}
}
