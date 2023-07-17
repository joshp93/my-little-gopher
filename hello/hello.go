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

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
