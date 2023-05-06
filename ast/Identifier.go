package ast

import "TroInterpreter/token"

type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // 标识符的值
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
