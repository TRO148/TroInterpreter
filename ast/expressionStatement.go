package ast

import "TroInterpreter/token"

type ExpressionStatement struct {
	Token      token.Token // 表达式语句的第一个token
	Expression Expression  // 表达式
}

//继承Statement

func (es *ExpressionStatement) statementNode() {}

// 继承Node
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
