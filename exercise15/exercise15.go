package exercise15

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
)

const PASSWORD = "melon"

// Password validation
func Run() {

	r := bufio.NewReader(os.Stdin)
	usrpass, err := util.GetInput(r, "What is the password? ")

	if err != nil {
		log.Fatalln(err)
	}

	if PASSWORD != usrpass {
		fmt.Println("I don't know you.")
	} else {
		fmt.Println("Welcome!")
	}

}
