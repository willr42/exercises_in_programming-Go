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
	bill, err := GetInput(r)
	if err != nil {
		log.Fatalln(err)
	}

	billCents, err := StringToCents(bill)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("What is the tip percentage? ")
	tipRate, err := GetInput(r)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(billCents)
	tipRateNum, err := strconv.Atoi(tipRate)
	if err != nil {
		log.Fatalln(err)
	}

	if tipRateNum > 100 || tipRateNum < 0 {
		log.Fatalln("must tip a valid percentage")
	}

	tipAmount := (billCents * tipRateNum) / 100
	tipStr := strconv.Itoa(tipAmount)
	fmt.Printf("The tip is $%s.%s\n", tipStr[:len(tipStr)-2], tipStr[len(tipStr)-2:])

	total := billCents + tipAmount
	totalStr := strconv.Itoa(total)
	fmt.Printf("The total is $%s.%s\n", totalStr[:len(totalStr)-2], totalStr[len(totalStr)-2:])
}

func GetInput(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", errors.New("couldn't read input")
	}
	finalLine := strings.Trim(line, "\n")
	return finalLine, nil
}

func StringToCents(s string) (int, error) {
	noDollar := strings.Replace(s, "$", "", -1)
	sArr := strings.Split(noDollar, ".")

	if len(sArr) > 2 {
		log.Fatalln("$$.CC")
	}

	// Only dollars
	if len(sArr) == 1 {
		dollars, err := strconv.Atoi(noDollar)
		if err != nil {
			return 0, errors.New("couldn't convert to int")
		}

		// to cents
		return dollars * 100, nil
	}

	centStr := sArr[1][0:2]
	finalStr := sArr[0] + centStr
	finalCents, err := strconv.Atoi(finalStr)
	if err != nil {
		return 0, errors.New("couldn't convert to int")
	}

	return finalCents, nil
}
