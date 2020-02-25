package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/francojposa/interpreter-in-go/01/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s. Welcome to Monkey\n", user.Username)
	fmt.Printf("LFG...\n")

	repl.Start(os.Stdin, os.Stdout)
}
