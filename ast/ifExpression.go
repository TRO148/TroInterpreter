package ast

import "TroInterpreter/token"

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out string
	for _, s := range bs.Statements {
		out += s.String()
	}
	return out
}

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
