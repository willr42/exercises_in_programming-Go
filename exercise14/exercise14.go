package exercise14

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Exercise
func Run() {

	r := bufio.NewReader(os.Stdin)
	amount, err := util.GetIntWithPrompt(r, "What is the order amount? ")
	if err != nil {
		log.Fatalln(err)
	}

	state, err := util.GetInput(r, "What is the state? ")
	if err != nil {
		log.Fatalln(err)
	}

	state = strings.ToLower(state)

	fmt.Println(printOutput(amount, state))
}

func printOutput(amount int, state string) string {

	if state == "wi" || state == "wisconsin" {
		taxRate := 5.5
		tax := int(math.Round(float64(amount) * taxRate))
		// as cents
		total := amount*100 + tax
		totalStr := util.FormatCents(total)
		subtotal := util.FormatCents(amount * 100)
		taxStr := util.FormatCents(tax)

		return fmt.Sprintf("The subtotal is $%s.\nThe tax is $%s.\nThe total is $%s.", subtotal, taxStr, totalStr)
	}
	totalStr := util.FormatCents(amount * 100)
	return fmt.Sprintf("The total is $%s.", totalStr)
}
