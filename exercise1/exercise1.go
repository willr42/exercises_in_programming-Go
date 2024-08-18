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

var customGreetings = map[string]string{"Brian": "Hello, %s! Have a great day.", "Susan": "Oh hey %s. How's your day looking?"}

func buildOutput(s string) string {

	greeting, ok := customGreetings[s]
	if !ok {
		return "Oh, a new face! Nice to meet you."
	}

	out := fmt.Sprintf(greeting, s)

	return out
}
