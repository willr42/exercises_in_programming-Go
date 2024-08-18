package exercise2

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"os"
)

// Counting the number of characters
func Run() {
	r := bufio.NewReader(os.Stdin)
	input := retryInput(r, "What is the input string? ")
	fmt.Printf("%s has %d characters.\n", input, len(input))
}

func retryInput(r *bufio.Reader, s string) string {
	for {
		input, err := util.GetInput(r, s)
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
