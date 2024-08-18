package exercise4

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
)

// Mad lib
func Run() {
	r := bufio.NewReader(os.Stdin)
	noun, err := util.GetInput(r, "Enter a noun: ")

	if err != nil {
		log.Fatalln(err)
	}
	verb, err := util.GetInput(r, "Enter a verb: ")

	if err != nil {
		log.Fatalln(err)
	}
	adjective, err := util.GetInput(r, "Enter a adjective: ")

	if err != nil {
		log.Fatalln(err)
	}
	adverb, err := util.GetInput(r, "Enter a adverb: ")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n", verb, adjective, noun, adverb)

}
