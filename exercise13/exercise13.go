package exercise13

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

	compoundOccurrence, err := util.GetIntWithPrompt(r, "What is the number of times the interest is compounded per year? ")

	if err != nil {
		log.Fatalln(err)
	}

	res := CompoundInterest(principal, rate, years, compoundOccurrence)
	total := util.FormatCents(res)
	original := util.FormatCents(principal)

	fmt.Printf("$%s invested at %.1f%% for %d years compounded %d times per year is $%s\n", original, rawRate, years, compoundOccurrence, total)

}

func CompoundInterest(principal int, rate float64, years int, compoundOccurrence int) int {
	baseRate := 1 + rate/float64(compoundOccurrence)
	yearsByCompound := float64(years * compoundOccurrence)
	total := float64(principal) * (math.Pow(baseRate, yearsByCompound))
	rounded := int(math.Round(total))
	return rounded
}
