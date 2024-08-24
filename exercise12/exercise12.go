package exercise12

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
)

// Exercise
func Run() {

	r := bufio.NewReader(os.Stdin)
	principal, err := util.GetFloatWithPrompt(r, "Enter the principal: ")

	if err != nil {
		log.Fatalln(err)
	}

	rate, err := util.GetFloatWithPrompt(r, "Enter the rate of interest: ")

	if err != nil {
		log.Fatalln(err)
	}

	rate = rate / 100

	years, err := util.GetIntWithPrompt(r, "Enter the number of years: ")

	if err != nil {
		log.Fatalln(err)
	}

	total := principal * (1 + rate*float64(years))

	fmt.Printf("After %d years at %f%%, the investment will be worth %f", years, rate, total)

}
