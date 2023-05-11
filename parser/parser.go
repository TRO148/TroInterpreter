package parser

import (
	"TroInterpreter/ast"
	"TroInterpreter/lexer"
	"TroInterpreter/token"
	"fmt"
	"strconv"
)

type Parser struct {
	l *lexer.Lexer //词法分析器

	curToken  token.Token //当前token
	peekToken token.Token //下一个token

	errors []string //错误

	prefixParseFns map[token.TypeToken]prefixParseFn //前缀解析函数映射
	infixParseFns  map[token.TypeToken]infixParseFn  //中缀解析函数映射
}

// 读取下一个token
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

// 报错类型
func (p *Parser) peekError(t token.TypeToken) {
	msg := fmt.Sprintf("类型得是 %s,却是 %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// 报错前缀
func (p *Parser) noPrefixParseFnError(t token.TypeToken) {
	msg := fmt.Sprintf("没有 %s 前缀解析函数", t)
	p.errors = append(p.errors, msg)
}

// 注册前缀解析函数
func (p *Parser) registerPrefix(tokenType token.TypeToken, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// 注册中缀解析函数
func (p *Parser) registerInfix(tokenType token.TypeToken, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// 获取当前token优先级
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

// 获取下一个token优先级
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
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
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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
	//分析表达式
	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}

// 分析return语句
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	//创建return语句节点
	stmt := &ast.ReturnStatement{Token: p.curToken}
	//跳过return
	p.nextToken()
	//分析表达式
	stmt.ReturnValue = p.parseExpression(LOWEST)
	return stmt
}

// 分析表达式语句
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken} //创建表达式语句节点
	stmt.Expression = p.parseExpression(LOWEST)         //分析表达式
	if p.peekToken.Type == token.SEMICOLON {            //如果下一个token是分号，跳过
		p.nextToken()
	}
	return stmt
}

// 分析表达式
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type] //获取前缀解析函数
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	// 执行前缀表达式解析函数
	leftExp := prefix()

	for p.peekToken.Type != token.SEMICOLON && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type] //获取中缀解析函数
		if infix == nil {
			return leftExp
		}
		p.nextToken() //到对应的中缀token再调用中缀解析函数
		leftExp = infix(leftExp)
	}

	return leftExp
}

//q: 解析器如何分清表达式，是否继续？例如if判断的时候1+1)

// 解析块语句
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}
	//跳过{
	p.nextToken()
	for p.curToken.Type != token.RBRACE && p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt) //将语句添加到语句数组中
		}
		//跳过;
		p.nextToken()
	}

	return block
}

// 分析if表达式
func (p *Parser) parseIfExpression() ast.Expression {
	exprssion := &ast.IfExpression{
		Token: p.curToken,
	}

	if !p.expectPeekAndNext(token.LPAREN) {
		return nil
	}

	p.nextToken()
	exprssion.Condition = p.parseExpression(LOWEST)

	//跳到)
	if !p.expectPeekAndNext(token.RPAREN) {
		return nil
	}
	//跳到{
	if !p.expectPeekAndNext(token.LBRACE) {
		return nil
	}

	//解析块语句
	exprssion.Consequence = p.parseBlockStatement()

	//跳到else
	if p.expectPeekAndNext(token.ELSE) {
		//跳到{
		if !p.expectPeekAndNext(token.LBRACE) {
			return nil
		}

		//解析块语句
		exprssion.Alternative = p.parseBlockStatement()
	}

	return exprssion
}

// 分析标识符
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} //创建标识符节点
}

// 分析整数
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken} //创建整数节点
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("无法解析 %q 为整数", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

// 分析布尔值
func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: token.TRUE == p.curToken.Type} //创建布尔值节点
}

// 分析分组表达式
func (p *Parser) parseGroupedExpression() ast.Expression {
	//过当前的token(
	p.nextToken()
	//解析下面的
	exp := p.parseExpression(LOWEST)

	if !p.expectPeekAndNext(token.RPAREN) {
		return nil
	}

	return exp
}

// 分析前缀表达式
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal, //前缀操作符
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

// 分析中缀表达式
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal, //中缀操作符
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	//读取两个token，初始化curToken和peekToken
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TypeToken]prefixParseFn)
	p.infixParseFns = make(map[token.TypeToken]infixParseFn)

	//注册前缀解析函数
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.NUMBER, p.parseIntegerLiteral)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)

	//注册中缀解析函数
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.ASTER, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)

	return p
}
