package exercise17

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
	"strings"
)

// Exercise
func Run() {

	r := bufio.NewReader(os.Stdin)
	alcohol, err := util.GetFloatWithPrompt(r, "Alcohol, in ounces: ")
	if err != nil {
		log.Fatalln(err)
	}

	weight, err := util.GetFloatWithPrompt(r, "Body weight, in pounds: ")
	if err != nil {
		log.Fatalln(err)
	}
	sex, err := util.GetInput(r, "(M)ale or (F)emale: ")
	if err != nil {
		log.Fatalln(err)
	}

	upper := strings.ToUpper(sex)
	if upper != "M" && upper != "F" {
		log.Fatalln("Must be (M) or (F)")
	}

	hours, err := util.GetFloatWithPrompt(r, "Hours since last drink: ")
	if err != nil {
		log.Fatalln(err)
	}

	res := calculateBAC(alcohol, weight, hours, sex)

	if res > 0.08 {
		log.Fatalf("Your BAC is %.2f\nIt is not legal for you to drive.\n", res)
	}
	fmt.Println("Drive away!")
}

func calculateBAC(alcohol, weight, hours float64, sex string) float64 {
	r := 0.0

	if sex == "M" {
		r = 0.73
	} else {
		r = 0.66
	}

	return (alcohol*5.14)/(weight*r) - (.015 * hours)
}
