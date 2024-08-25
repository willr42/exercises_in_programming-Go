package exercise15

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/term"
)

type Creds struct {
	user     string
	password string
}

var Users = []Creds{{user: "alice", password: "wonderland"}, {user: "bob", password: "thebuilder"}}

// Password validation
func Run() {

	r := bufio.NewReader(os.Stdin)
	user, err := util.GetInput(r, "What is your username? ")
	if err != nil {
		log.Fatalln(err)
	}

	found := false
	targetPass := ""

	for _, v := range Users {
		if user == v.user {
			found = true
			targetPass = v.password
		}
	}

	if !found {
		log.Fatalf("User %s not found.\n", user)
	}

	fmt.Print("What is your password? ")
	bytepass, err := term.ReadPassword(syscall.Stdin)

	if err != nil {
		log.Fatalln(err)
	}

	usrpass := string(bytepass)

	if targetPass != usrpass {
		fmt.Println("\nI don't know you.")
	} else {
		fmt.Println("\nWelcome!")
	}

}
