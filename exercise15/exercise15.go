package exercise15

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
)

const PASSWORD = "melon"

// Password validation
func Run() {

	fmt.Print("What is the password? ")
	bytepass, err := term.ReadPassword(syscall.Stdin)

	if err != nil {
		log.Fatalln(err)
	}

	usrpass := string(bytepass)

	if PASSWORD != usrpass {
		fmt.Println("I don't know you.")
	} else {
		fmt.Println("Welcome!")
	}

}
