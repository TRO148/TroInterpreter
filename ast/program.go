package ast

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
