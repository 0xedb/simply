package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/thebashshell/simply/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello, %s\n Welcome to the SIMPLY PROGRAMMING LANGUAGE:\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
