package token

type TokenType string

type Token struct {
	Type    TokenType // 区分不同类型词法单元
	Literal string
}

const (
	ILLEGAL = "ILLEAGL" // 未知词法单元/字符
	EOF     = "EOF"     // 文件结尾

	// 标识符 字面两
	IDENT  = "IDENT" // add, foobar, x, y
	INT    = "INT"   // 123
	ASSIGN = "="
	PLUS   = "+"

	// 运算符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
