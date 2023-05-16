package lexer

import (
	"TroInterpreter/token"
)

type Lexer struct {
	input        string //输入
	position     int    //字符在输入中的位置
	readPosition int    //下一个字符位置
	ch           byte   //当前字符
}

func (l *Lexer) readChar() {
	//读取指针更新
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	//预读取
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	//读取标识符
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumberIdentifier() string {
	//读取数字
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	//跳过空白字符
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		//读取指针更新
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	//读取字符串
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace() //跳过空白字符

	//根据字符返回不同的token，并且读取指针前移
	var tok token.Token
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch) //创建一个新的token
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)} //创建一个新的token
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch) //创建一个新的token
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = token.Token{Type: token.BANG, Literal: string(l.ch)} //创建一个新的token
		}

		//分隔符
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
	case '[':
		tok = token.Token{Type: token.LBRACKET, Literal: string(l.ch)}
	case ']':
		tok = token.Token{Type: token.RBRACKET, Literal: string(l.ch)}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	case '"':
		tok = token.Token{Type: token.STRING, Literal: l.readString()}
		//运算符
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: string(l.ch)}
	case '*':
		tok = token.Token{Type: token.ASTER, Literal: string(l.ch)}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: string(l.ch)}
	case '<':
		tok = token.Token{Type: token.LT, Literal: string(l.ch)}
	case '>':
		tok = token.Token{Type: token.GT, Literal: string(l.ch)}

		//标识符
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()          //读取标识符
			tok.Type = token.LookupIdent(tok.Literal) //根据标识符返回对应的token类型
			return tok                                //跳过readChar()
		} else if isNumber(l.ch) {
			tok.Type = token.NUMBER
			tok.Literal = l.readNumberIdentifier()
			return tok
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = string(l.ch)
		}
	}
	l.readChar()
	return tok
}
