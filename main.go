package main

import (
	"TroInterpreter/repl"
	"os"
)

func main() {
	println("欢迎使用Tro，调用help()查看更多信息")
	repl.Start(os.Stdin, os.Stdout)
}
