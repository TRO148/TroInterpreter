package repl

import (
	"TroInterpreter/evaluator"
	"TroInterpreter/lexer"
	"TroInterpreter/object"
	"TroInterpreter/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func printParserError(out io.Writer, err []string) {
	io.WriteString(out, "解析错误:\n")
	for _, msg := range err {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func Start(in io.Reader, out io.Writer) {
	//创建输入输出流
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Fprintf(out, PROMPT)
		//读取输入
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		//解析输入
		line := scanner.Text()
		//创建词法分析器
		l := lexer.New(line)
		//创建语法分析器
		p := parser.New(l)

		//解析程序
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserError(out, p.Errors())
			continue
		}

		//求值器
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
