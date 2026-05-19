package main

import (
	"fmt"
	"interpreter/src/repl"
	"os"
)

func main() {
	fmt.Println("Welcome to Monkey interpreter, type commands:")
	repl.Start(os.Stdin, os.Stdout)
}
