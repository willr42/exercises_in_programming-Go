package exercise16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Exercise
func Run() {

	r := bufio.NewReader(os.Stdin)
	fmt.Print("What is your age? ")

	age := 0

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading string")
		}
		line = strings.TrimSpace(line)
		intInp, err := strconv.Atoi(line)

		if err == nil {
			age = intInp
			break
		}

		fmt.Print("Must be a valid integer.\n")

		fmt.Print("Your age? ")
	}

	if age < 16 {
		fmt.Println("You are not old enough to legally drive.")
	} else {
		fmt.Println("You are old enough to legally drive.")

	}

}
