package ast

import (
	"TroInterpreter/token"
	"bytes"
)

// 中缀表达式
type InfixExpression struct {
	Token    token.Token // 操作符
	Left     Expression  // 左表达式
	Operator string      // 操作符
	Right    Expression  // 右表达式
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer // 字节缓冲区

	//读取
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
