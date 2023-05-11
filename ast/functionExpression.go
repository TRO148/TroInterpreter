package ast

import "TroInterpreter/token"

// 函数表达式
type FunctionExpression struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fe *FunctionExpression) expressionNode() {}
func (fe *FunctionExpression) TokenLiteral() string {
	return fe.Token.Literal
}
func (fe *FunctionExpression) String() string {
	var out string

	out += "func("
	for i, p := range fe.Parameters {
		if i != 0 {
			out += ", "
		}
		out += p.String()
	}
	out += ") "
	out += fe.Body.String()

	return out
}
