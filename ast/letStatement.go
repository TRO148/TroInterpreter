package ast

import (
	"TroInterpreter/token"
)

// 实现Statement接口，能够将LetStatement添加到Program中
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // 标识符
	Value Expression  // 表达式
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
