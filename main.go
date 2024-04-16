package main

import (
	"chirp/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Chirp programming language\n", user.Username)
	fmt.Printf("Start Chirping some commands & see us generate tokens from the code!!!\n")
	repl.Start(os.Stdin, os.Stdout)
}
