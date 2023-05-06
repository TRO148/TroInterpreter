package repl

import (
	"TroInterpreter/lexer"
	"TroInterpreter/token"
	"bufio"
	"fmt"
	"io"
	"time"
)

const PROMPT = " >>  "

func Start(in io.Reader, out io.Writer) {
	//创建输入输出流
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("# " + time.TimeOnly + PROMPT)
		//读取输入
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		//解析输入
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			//输出解析结果
			fmt.Printf("%+v\n", tok)
		}
	}
}
