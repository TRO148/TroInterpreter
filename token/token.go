package token

type TypeToken string

type Token struct {
	Type    TypeToken //token类型
	Literal string    //token字面量
}

var keywords = map[string]TypeToken{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TypeToken {
	//根据标识符返回对应的token类型
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	// 特殊标记
	ILLEGAL = "ILLEGAL" //非法字符，表示遇到未知的词法单元
	EOF     = "EOF"     //文件结束，通知语法分析器停机
	// 运算符
	ASSIGN = "="  //赋值
	PLUS   = "+"  //加法
	MINUS  = "-"  //减法
	BANG   = "!"  //感叹号
	ASTER  = "*"  //乘法
	SLASH  = "/"  //除法
	LT     = "<"  //小于
	GT     = ">"  //大于
	EQ     = "==" //等于
	NOT_EQ = "!=" //不等于
	// 分隔符
	COMMA     = "," //逗号
	SEMICOLON = ";" //分号
	LPAREN    = "(" //左括号
	RPAREN    = ")" //右括号
	LBRACE    = "{" //左花括号
	RBRACE    = "}" //右花括号
	// 标识符
	IDENT = "IDENT" //标识符
	// 关键字
	FUNCTION = "FUNCTION" //函数
	LET      = "LET"      //变量声明
	TRUE     = "TRUE"     //真
	FALSE    = "FALSE"    //假
	IF       = "IF"       //if
	ELSE     = "ELSE"     //else
	RETURN   = "RETURN"   //return
	NUMBER   = "NUMBER"   //数字
	STRING   = "STRING"   //字符串
)
