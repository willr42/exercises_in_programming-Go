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
	fmt.Print("What is the quote? ")
	quote, err := util.GetInput(r)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Who said it? ")
	author, err := util.GetInput(r)

	if err != nil {
		log.Fatalln(err)
	}

	final := formatOutput(quote, author)

	fmt.Println(final)
}

func formatOutput(quote string, author string) string {
	return author + " says, \"" + quote + "\""
}
