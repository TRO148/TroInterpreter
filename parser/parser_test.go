package parser

import (
	"TroInterpreter/ast"
	"TroInterpreter/lexer"
	"fmt"
	"testing"
)

func TestStatements(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.value) {
			return
		}
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

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	// 类型断言
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok { //判断是否为IntegerLiteral
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) { //判断整数值是否正确
		t.Errorf("integ.TokenLiteral() not %d. got=%s", value, integ.TokenLiteral())
		return false
	}

	if integ.Value != value { //判断整数值是否正确
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	return true
}
