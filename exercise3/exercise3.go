package exercise3

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
)

// Printing quotes
func Run() {
	r := bufio.NewReader(os.Stdin)
	quote, err := util.GetInput(r, "What is the quote? ")

	if err != nil {
		log.Fatalln(err)
	}
	author, err := util.GetInput(r, "Who said it? ")

	if err != nil {
		log.Fatalln(err)
	}

	final := formatOutput(quote, author)

	fmt.Println(final)
}

func formatOutput(quote string, author string) string {
	return author + " says, \"" + quote + "\""
}
