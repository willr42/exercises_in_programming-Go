package exercise7

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
)

var ftToM = 0.09290304

// Exercise
func Run() {

	r := bufio.NewReader(os.Stdin)

	length, err := util.GetIntWithPrompt(r, "What is the length of the room in feet? ")

	if err != nil {
		log.Fatalln(err)
	}

	width, err := util.GetIntWithPrompt(r, "What is the width of the room in feet? ")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("You entered dimensions of %d feet by %d feet.\n", length, width)

	area := length * width
	fmt.Println("The area is")

	fmt.Printf("%d square feet\n", area)
	inMeters := float64(area) * ftToM
	fmt.Printf("%.3f square meters\n", inMeters)

}
