package ast

import "TroInterpreter/token"

type IfExpression struct {
	Token       token.Token // if
	Condition   Expression
	Consequence *BlockStatement //结果
	Alternative *BlockStatement //第二个结果
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out string

	out += "if"
	out += ie.Condition.String()
	out += " "
	out += ie.Consequence.String()

	if ie.Alternative != nil {
		out += "else "
		out += ie.Alternative.String()
	}

	return out
}
