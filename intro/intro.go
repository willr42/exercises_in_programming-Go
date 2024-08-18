package intro

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// TODO: GUI where function runs live
func CalculateTip() {
	r := bufio.NewReader(os.Stdin)

	billCents := getBillCents(r, "What is the bill? ")
	tipRate := getTipRate(r, "What is the tip percentage? ")

	tipAmount := (billCents * tipRate) / 100
	tipStr := formatCents(tipAmount)
	fmt.Printf("The tip is $%s\n", tipStr)

	total := billCents + tipAmount
	totalStr := formatCents(total)
	fmt.Printf("The total is $%s\n", totalStr)
}

func getBillCents(r *bufio.Reader, s string) int {
	for {
		bill, err := util.GetInput(r, s)
		if err != nil {
			fmt.Print("Invalid input. Try again: ")
			continue
		}

		billCents, err := stringToCents(bill)
		if err == nil {
			return billCents
		}
		fmt.Print("Invalid input. Try again: ")
	}
}

func getTipRate(r *bufio.Reader, s string) int {
	for {
		tipRate, err := util.GetInput(r, s)
		if err != nil {
			fmt.Print("Invalid input. Try again: ")
			continue
		}

		tipRateNum, err := strconv.Atoi(tipRate)
		if err != nil {
			fmt.Print("Invalid input. Try again: ")
			continue
		}

		if tipRateNum > 100 || tipRateNum < 0 {
			fmt.Print("Tip rate must be between 0 & 100. Try again: ")
			continue
		}

		return tipRateNum
	}
}

func stringToCents(s string) (int, error) {
	noDollar := strings.Replace(s, "$", "", -1)
	parts := strings.Split(noDollar, ".")

	if len(parts) > 2 {
		log.Fatalln("$$.CC")
	}

	dollars := 0
	cents := 0

	if len(parts) > 0 {
		var err error
		dollars, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, fmt.Errorf("couldn't convert dollars to int: %w", err)
		}
	}

	// Cents exists
	if len(parts) == 2 {
		if len(parts[1]) > 2 {
			parts[1] = parts[1][:2]
		}
		var err error
		cents, err = strconv.Atoi(parts[1])
		if err != nil {

			return 0, fmt.Errorf("couldn't convert dollars to int: %w", err)
		}

	}
	return dollars*100 + cents, nil

}

func formatCents(cents int) string {
	dollars := cents / 100
	remainingCents := cents % 100
	return fmt.Sprintf("%d.%02d", dollars, remainingCents)
}
