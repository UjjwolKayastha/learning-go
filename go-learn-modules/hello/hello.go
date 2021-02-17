package main

import (
	"fmt"
	"log"

	"al.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"AL", "UJJWOL", "Kayastha"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// message := greetings.Hello("UJJWOL")
	fmt.Println(message)
}
