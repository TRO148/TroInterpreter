package repl

import (
	"TroInterpreter/lexer"
	"TroInterpreter/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = " >>  "

func printParserError(out io.Writer, err []string) {
	io.WriteString(out, "解析错误:\n")
	for _, msg := range err {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func Start(in io.Reader, out io.Writer) {
	//创建输入输出流
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		//读取输入
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		//解析输入
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserError(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}
