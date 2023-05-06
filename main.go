package main

import (
	"TroInterpreter/repl"
	"os"
)

func main() {
	println("Here is Tro！")
	repl.Start(os.Stdin, os.Stdout)
}
