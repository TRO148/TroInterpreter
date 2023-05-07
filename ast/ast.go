package ast

import (
	"TroInterpreter/token"
	"bytes"
)

// ast节点都应实现Node接口
type Node interface {
	TokenLiteral() string
	String() string
}

// 语句
type Statement interface {
	Node
	statementNode()
}

// 表达式
type Expression interface {
	Node
	expressionNode()
}

// 程序
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 { // 语句不为空
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
func (p *Program) String() string {
	var out bytes.Buffer // 字节缓冲区

	//读取
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

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
