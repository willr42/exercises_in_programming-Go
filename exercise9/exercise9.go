package exercise9

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
)

var gallonToFeet = 350

// Exercise
func Run() {

	r := bufio.NewReader(os.Stdin)
	l, err := util.GetIntWithPrompt(r, "Room length: ")

	if err != nil {
		log.Fatalln(err)
	}

	w, err := util.GetIntWithPrompt(r, "Room width: ")

	if err != nil {
		log.Fatalln(err)
	}

	totalArea := w * l

	var gallonsRequired = 0
	areaCounter := totalArea

	for {
		if areaCounter < gallonToFeet {
			// Always round up
			gallonsRequired++
			break
		}
		areaCounter = areaCounter - gallonToFeet
		gallonsRequired++
	}

	fmt.Printf("You will need to purchase %d gallons of paint to cover %d square feet.\n", gallonsRequired, totalArea)
}
