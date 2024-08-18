package exercise6

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
	"time"
)

// Retirement calculator
func Run() {

	r := bufio.NewReader(os.Stdin)
	ageStr, err := util.GetInput(r, "What is your current age? ")

	if err != nil {
		log.Fatalln(err)
	}
	age, err := util.StrToInt(ageStr)

	if err != nil {
		log.Fatalln(err)
	}

	retirementAgeStr, err := util.GetInput(r, "At what age would you like to retire? ")

	if err != nil {
		log.Fatalln(err)
	}

	retirementAge, err := util.StrToInt(retirementAgeStr)

	if err != nil {
		log.Fatalln(err)
	}

	if retirementAge < 0 {
		fmt.Println("You're already retirement-ready!")
		os.Exit(0)
	}

	yearsLeft := retirementAge - age

	fmt.Printf("You have %d years left until you can retire.\n", yearsLeft)

	now := time.Now()
	fmt.Printf("It's %d, so you can retire in %d.\n", now.Year(), now.Year()+yearsLeft)
}
