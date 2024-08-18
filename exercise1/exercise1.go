package exercise1

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
)

func Run() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	input, err := util.GetInput(r)
	if err != nil {
		log.Fatalln(err)
	}

	final := buildOutput(input)
	fmt.Println(final)
}

func buildOutput(s string) string {

	out := fmt.Sprintf("Hello, %s, nice to meet you!", s)

	return out
}
