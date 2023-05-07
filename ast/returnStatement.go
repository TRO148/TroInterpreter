package ast

import (
	"TroInterpreter/token"
	"bytes"
)

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

// 继承statement
func (rs *ReturnStatement) statementNode() {}

// 继承node
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
