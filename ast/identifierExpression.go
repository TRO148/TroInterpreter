package ast

import "TroInterpreter/token"

// 标识符
type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // 标识符的值
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}
