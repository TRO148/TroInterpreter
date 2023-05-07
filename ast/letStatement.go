package ast

import (
	"TroInterpreter/token"
	"bytes"
)

// 实现Statement接口，能够将LetStatement添加到Program中
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // 标识符
	Value Expression  // 表达式
}

// 继承statement
func (ls *LetStatement) statementNode() {}

// 继承node
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")

	out.WriteString(ls.Name.String())

	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
