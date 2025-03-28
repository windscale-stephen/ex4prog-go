// Package tipcalc contains functions used to build a tip calculator.
package tipcalc

import (
	"fmt"
	"io"
	"log"
	"math"
)

// DisplayPrompt writes the given prompt to the given io.Writer.
// If it's unable to successfully write the prompt it displays an error message
// and exits the program.
func DisplayPrompt(w io.Writer, prompt string) error {
	_, err := fmt.Fprintf(w, prompt)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// ReadFPNumber reads and returns a floating point number from the given
// io.Reader.
// It discards the input line and returns an error if it isn't a valid
// floating-point number.
func ReadFPNumber(r io.Reader) (float64, error) {
	var fp float64
	_, err := fmt.Fscanln(r, &fp)
	if err != nil {
		var discard string
		fmt.Fscanln(r, &discard)
		return 0.0, fmt.Errorf("must be a valid decimal number")
	}
	return fp, nil
}

// Tip calculates the tip for the given bill and tip percentage.
// It returns an error if the bill is less than or equal to zero or if the tip
// is less than zero.
func Tip(bill, tipPct float64) (float64, error) {
	if bill <= 0.0 {
		return 0.0, fmt.Errorf("bill %f must be greater than zero", bill)
	}
	if tipPct < 0.0 {
		return 0.0, fmt.Errorf("tip percentage %f must be greater than or equal to zero", tipPct)
	}
	tip := (bill * tipPct) / 100
	tip = math.Ceil(tip*100) / 100 // round up to 2 d.p.
	return tip, nil
}

// Total calculates the total amount given the bill and tip.
// It returns an error if the bill is less than or equal to zero or if the tip
// is less than zero.
func Total(bill, tip float64) (float64, error) {
	if bill <= 0.0 {
		return 0.0, fmt.Errorf("bill %f must be greater than zero", bill)
	}
	if tip < 0.0 {
		return 0.0, fmt.Errorf("tip %f must be greater than or equal to zero", tip)
	}
	return bill + tip, nil
}
