package exercise2

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"os"
)

func Run() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("What is the input string? ")
	input := retryInput(r)
	fmt.Printf("%s has %d characters.\n", input, len(input))
}

func retryInput(r *bufio.Reader) string {
	for {
		input, err := util.GetInput(r)
		if len(input) == 0 {
			fmt.Println("Sorry, there must be input.")
			fmt.Print("What is the input string? ")
			continue
		}

		if err == nil {
			return input
		}
	}
}
