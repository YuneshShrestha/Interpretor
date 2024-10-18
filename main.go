package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/YuneshShrestha/Interpretor/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello " + user.Username + "!")
	fmt.Println("Feel free to type in commands")
	fmt.Println("To exit the REPL, press Ctrl+C")
	repl.Start(os.Stdin, os.Stdout)
}
