// Sayhello prompts for a name and then prints a greeting.
package main

import (
	"log"
	"os"

	"github.com/windscale-stephen/ex4prog_go/ch02/sayhello/internal/sayhello"
)

func main() {
	// TODO: Currently have to enter EOF after entering name to display greeting.
	sayhello.DisplayPrompt(os.Stdout, "What is your name? ")
	name, err := sayhello.ReadName(os.Stdin)
	if err != nil {
		log.Fatal("must enter a valid name")
	}
	greeting := sayhello.MakeGreeting(name)
	sayhello.DisplayPrompt(os.Stdout, greeting)
}
