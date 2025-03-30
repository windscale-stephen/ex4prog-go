package sayhello

import (
	"fmt"
	"io"
	"log"
)

// DisplayPrompt writes the given prompt to the given io.Writer.
// If it's unable to successfully write the prompt it displays an error message
// and exits the program.
func DisplayPrompt(w io.Writer, prompt string) error {
	_, err := fmt.Fprint(w, prompt)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// MakeGreeting returns a greeting based on the supplied name.
func MakeGreeting(name string) string {
	return "Hello, " + name + ", nice to meet you!\n"
}

// ReadName reads and returns a name from the given io.Reader.
// It discards the input line and returns an error if it isn't a valid string.
func ReadName(r io.Reader) (string, error) {
	var name string
	sep := ""
	for {
		var input string
		_, err := fmt.Fscan(r, &input)
		if err == io.EOF {
			break
		}
		if err != nil {
			var discard string
			fmt.Fscanln(r, &discard)
			// TODO: wrap message from err with own message.
			return "", fmt.Errorf("must be a valid string")
		}
		name = name + sep + input
		sep = " "
	}
	return name, nil
}
