package lexer

import "monkeyLang/token"

type Lexer struct {
	input        string
	position     int  // 当前位置
	readPosition int  // 读取位置，指向当前字符之后一个字符
	ch           byte // 当前查看字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // 初始化
	return l
}

func (l *Lexer) readChar() {
	// 如果达到input末尾，将其置为0,
	// 否则置为下一个字符
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	// 两个指针向后移动
	l.position = l.readPosition
	l.readPosition++
}

// 查看当前的字符，根据具体字符，返回对应词法单元,然后当前指针移动,保证下次调用，已经更新
func (l *Lexer) NextToken() token.Token {

	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
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

// 返回一个Token结构体
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
