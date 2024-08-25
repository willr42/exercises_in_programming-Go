package exercise12

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
	principal, err := util.GetCentsWithPrompt(r, "Enter the principal: ")

	if err != nil {
		log.Fatalln(err)
	}

	rawRate, err := util.GetFloatWithPrompt(r, "Enter the rate of interest: ")

	if err != nil {
		log.Fatalln(err)
	}

	rate := rawRate / 100

	years, err := util.GetIntWithPrompt(r, "Enter the number of years: ")

	if err != nil {
		log.Fatalln(err)
	}

	res := float64(principal) * (1 + rate*float64(years))
	total := util.FormatCents(int(math.Round(res)))

	fmt.Printf("After %d years at %.1f%%, the investment will be worth $%s\n", years, rawRate, total)

}
