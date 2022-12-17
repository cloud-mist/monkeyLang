package token

// ----------------------------------------
// 将要解析的code

// let five = 5;
// let ten = 10;
//
// let add = fn(x, y) {
//    x + y;
// };
//
// let result = add(five, ten);

// ----------------------------------------

// Token data structure
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// const: possibe TokenTypes
const (
	ILLEGAL = "ILLEGAL" // a token/char  dont know
	EOF     = "EOF"     // end of line

	IDENT = "IDENT"
	INT   = "INT"

	// 操作符
	ASSIGN = "="
	PLUS   = "+"

	// 分割标点
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键词
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
