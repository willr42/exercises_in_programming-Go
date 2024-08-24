package exercise11

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"math"
	"os"
)

// Exercise
func Run() {

	r := bufio.NewReader(os.Stdin)
	euro, err := util.GetFloatWithPrompt(r, "How many euros are you exchanging? ")

	if err != nil {
		log.Fatalln(err)
	}

	exchangeRate, err := util.GetFloatWithPrompt(r, "What is the exchange rate? ")

	if err != nil {
		log.Fatalln(err)
	}

	res := convertCurrency(euro, exchangeRate)

	fmt.Printf("%.2f euros at an exchange rate of %.2f is %.2f U.S. Dollars.\n", euro, exchangeRate, res)

}

func convertCurrency(from float64, rate float64) float64 {
	converted := from * (rate / 100)
	return math.Round(converted*100) / 100
}
