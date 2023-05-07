package ast

import "TroInterpreter/token"

type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // 标识符的值
}

// 继承expression
func (i *Identifier) expressionNode() {}

// 继承node
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}
