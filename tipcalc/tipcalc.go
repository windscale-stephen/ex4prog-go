// Tipcalc prompts for a bill amount and tip percentage and calculates the tip
// and total amount.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/windscale-stephen/ex4prog_go/tipcalc/internal/tipcalc"
)

func main() {
	log.SetFlags(0) // Suppress printing data & time prefix to log messages.
	tipcalc.DisplayPrompt(os.Stdout, "What is the bill? $")
	bill, err := tipcalc.ReadFPNumber(os.Stdin)
	if err != nil {
		log.Fatal("bill must be a valid number")
	}
	tipcalc.DisplayPrompt(os.Stdout, "What is the tip percentage? ")
	tipPct, err := tipcalc.ReadFPNumber(os.Stdin)
	if err != nil {
		log.Fatal("tip percentage must be a valid number")
	}
	tip, err := tipcalc.Tip(bill, tipPct)
	if err != nil {
		log.Fatal(err)
	}
	total, err := tipcalc.Total(bill, tip)
	if err != nil {
		log.Fatal(err)
	}
	tipcalc.DisplayPrompt(os.Stdout, fmt.Sprintf("The tip is $%.2f\n", tip))
	tipcalc.DisplayPrompt(os.Stdout, fmt.Sprintf("The total is $%.2f\n",
		total))
}
