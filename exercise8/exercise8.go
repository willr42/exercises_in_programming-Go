package exercise8

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
	people, err := util.GetIntWithPrompt(r, "How many people? ")

	if err != nil {
		log.Fatalln(err)
	}

	pizza, err := util.GetIntWithPrompt(r, "How many pizzas? ")

	if err != nil {
		log.Fatalln(err)
	}

	slices, err := util.GetIntWithPrompt(r, "Slices per pizza? ")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%d people with %d pizzas\n", people, pizza)

	sliceTotal := pizza * slices
	perPerson := sliceTotal / people
	remainder := sliceTotal % people

	var personPieceStr = "piece"

	if perPerson > 1 {
		personPieceStr = personPieceStr + "s"
	}

	fmt.Printf("Each person gets %d %s of pizza.\n", perPerson, personPieceStr)

	var remainderPieceStr = "piece"

	if remainder > 1 {
		remainderPieceStr = remainderPieceStr + "s"
	}

	var isTense = "is"

	if remainder > 1 {
		isTense = "are"
	}

	fmt.Printf("There %s %d leftover %s.\n", isTense, remainder, remainderPieceStr)

}
