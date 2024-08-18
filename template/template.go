package template

import (
	"bufio"
	"exercises-in-programming/util"
	"log"
	"os"
)

// Exercise
func Run() {

	r := bufio.NewReader(os.Stdin)
	_, err := util.GetInput(r, "Prompt: ")

	if err != nil {
		log.Fatalln(err)
	}
}
