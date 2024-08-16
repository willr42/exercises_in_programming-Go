package chapter1

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// - tip
// - total amount
// - prompt for bill & tip
// - compute tip
// - display tip & final total
func CalculateTip() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("What is the bill? ")
	billCents := getBillCents(r)

	fmt.Print("What is the tip percentage? ")
	tipRate := getTipRate(r)

	tipAmount := (billCents * tipRate) / 100
	tipStr := formatCents(tipAmount)
	fmt.Printf("The tip is $%s\n", tipStr)

	total := billCents + tipAmount
	totalStr := formatCents(total)
	fmt.Printf("The total is $%s\n", totalStr)
}

func getBillCents(r *bufio.Reader) int {
	for {
		bill, err := getInput(r)
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

func getTipRate(r *bufio.Reader) int {
	for {
		tipRate, err := getInput(r)
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

func getInput(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", errors.New("couldn't read input")
	}
	finalLine := strings.Trim(line, "\n")
	return finalLine, nil
}

func stringToCents(s string) (int, error) {
	noDollar := strings.Replace(s, "$", "", -1)
	parts := strings.Split(noDollar, ".")

	if len(parts) > 2 {
		log.Fatalln("$$.CC")
	}

	// Only dollars
	if len(parts) == 1 {
		dollars, err := strconv.Atoi(noDollar)
		if err != nil {
			return 0, errors.New("couldn't convert to int")
		}

		// to cents
		return dollars * 100, nil
	}

	centStr := parts[1][:2]
	fmt.Print(centStr)
	finalStr := parts[0] + centStr
	finalCents, err := strconv.Atoi(finalStr)
	if err != nil {
		return 0, errors.New("couldn't convert to int")
	}

	return finalCents, nil
}

func formatCents(cents int) string {
	dollars := cents / 100
	remainingCents := cents % 100
	return fmt.Sprintf("%d.%02d", dollars, remainingCents)
}
