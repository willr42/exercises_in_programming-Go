package exercise9

import (
	"bufio"
	"exercises-in-programming/util"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

var gallonToFeet = 350

// Exercise
func Run() {
	roundPtr := flag.Bool("round", false, "is shape round")
	flag.Parse()
	r := bufio.NewReader(os.Stdin)

	var totalArea = 0

	if *roundPtr {
		totalArea = int(circleArea(r))
	} else {
		totalArea = rectArea(r)
	}

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

func rectArea(r *bufio.Reader) int {
	l, err := util.GetIntWithPrompt(r, "Room length: ")

	if err != nil {
		log.Fatalln(err)
	}

	w, err := util.GetIntWithPrompt(r, "Room width: ")

	if err != nil {
		log.Fatalln(err)
	}

	return w * l
}

func circleArea(r *bufio.Reader) float64 {
	d, err := util.GetFloatWithPrompt(r, "Room diameter: ")

	if err != nil {
		log.Fatalln(err)
	}

	area := math.Pi * (math.Pow((d / 2), 2))

	return area
}
