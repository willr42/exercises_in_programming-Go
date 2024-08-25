package exercise15

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

type Creds struct {
	user     string
	password []byte
}

func DummyData() []Creds {
	nonEncrypt := []Creds{{user: "alice", password: []byte("wonderland")}, {user: "bob", password: []byte("thebuilder")}}
	encrypted := []Creds{}

	for _, v := range nonEncrypt {
		hash, err := bcrypt.GenerateFromPassword(v.password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		newUser := Creds{user: v.user, password: hash}
		encrypted = append(encrypted, newUser)
	}

	return encrypted

}

// Password validation
func Run() {

	Users := DummyData()

	r := bufio.NewReader(os.Stdin)
	user, err := util.GetInput(r, "What is your username? ")
	if err != nil {
		log.Fatalln(err)
	}

	found := false
	targetPass := []byte("")

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

	result := bcrypt.CompareHashAndPassword(targetPass, bytepass)

	if result != nil {
		fmt.Println("\nI don't know you.")
	} else {
		fmt.Println("\nWelcome!")
	}

}
