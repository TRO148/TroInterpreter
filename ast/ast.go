package ast

/*
接口，所有的节点都要实现这个接口
*/

type Node interface {
	TokenLiteral() string

	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
