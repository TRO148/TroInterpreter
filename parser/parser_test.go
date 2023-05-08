package parser

import (
	"TroInterpreter/ast"
	"TroInterpreter/lexer"
	"testing"
)

func TestStatements(t *testing.T) {
	input := `5;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != 5 {
		t.Errorf("ident.Value not %d. got=%d", 5, ident.Value)
	}
	if ident.TokenLiteral() != "5" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "5", ident.TokenLiteral())
	}
}

// 辅助函数

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors
	if len(errors) == 0 { //如果没有错误，直接返回
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors { //打印错误
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" { //判断是否为let语句
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	// 类型断言
	letStmt, ok := s.(*ast.LetStatement)
	if !ok { //判断是否为LetStatement
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name { //判断标识符是否为name
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name { //判断标识符是否为name
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
