package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/x-color/monkey/exec"
)

func main() {
	if len(os.Args) == 1 {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programing language!\n",
			user.Username)
		fmt.Println("Feel free to type in commands")
		exec.Repl(os.Stdin, os.Stdout)
	} else {
		fileName := os.Args[1]
		exec.ExecFile(fileName, os.Stdout)
	}
}
