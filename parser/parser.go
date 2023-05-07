package parser

import (
	"TroInterpreter/ast"
	"TroInterpreter/lexer"
	"TroInterpreter/token"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer //词法分析器

	curToken  token.Token //当前token
	peekToken token.Token //下一个token

	errors []string //错误
}

// 通过语法分析器，我们可以读取两个token，curToken和peekToken
func (p *Parser) nextToken() {
	//读取指针更新
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// 读取下一个token，判断是不是需要的
func (p *Parser) expectPeekAndNext(t token.TypeToken) bool {
	if p.peekToken.Type == t { //如果下一个token是t类型
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// 报错
func (p *Parser) peekError(t token.TypeToken) {
	msg := fmt.Sprintf("类型得是 %s,却是 %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// ParseProgram 分析程序
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}              //创建一个Program节点
	program.Statements = []ast.Statement{} //初始化语句数组

	// 读取每个token，直到遇到EOF
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt) //将语句添加到语句数组中
		}
		p.nextToken()
	}

	return program
}

// 分析语句，用来导向每一个具体的语句分析函数
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET: //let语句
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// 分析let语句
func (p *Parser) parseLetStatement() *ast.LetStatement {
	//创建let语句节点
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeekAndNext(token.IDENT) { //判断下一个token是否为IDENT
		return nil
	}
	//创建标识符节点
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} //创建标识符节点
	if !p.expectPeekAndNext(token.ASSIGN) {                                   //判断下一个token是否为ASSIGN
		return nil
	}
	//跳过=
	p.nextToken()
	//跳过表达式
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	return stmt
}

// 分析return语句
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	//创建return语句节点
	stmt := &ast.ReturnStatement{Token: p.curToken}
	//跳过return
	p.nextToken()
	//跳过表达式
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	return stmt
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	//读取两个token，初始化curToken和peekToken
	p.nextToken()
	p.nextToken()

	return p
}
