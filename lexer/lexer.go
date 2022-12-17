package lexer

import "monkeyLang/token"

type Lexer struct {
	input string //

	// 两个指针
	position     int // current char
	readPosition int // current's next char

	ch byte // current char under examination

}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 如果下个位置超界，就变成0,如果没有，就将其置为下个位置的值.然后当前位置指向下个位置
func (l *Lexer) readChar() {
	//{{{
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition // 下个位置
	l.readPosition++            // 下个位置+1
	// }}}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
