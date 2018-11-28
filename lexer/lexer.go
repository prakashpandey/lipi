package lexer

import (
	"github.com/prakashpandey/lipi/token"
)

// Lexer is data structure holding lexing details
type Lexer struct {
	input        string
	position     int  // position of current char
	readPosition int  // next character position to read
	ch           byte // current char in examination
}

// New returns new object of Lexer data structure
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // so that our lexer is in fully working state before anyone calls l.NextToken()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 is the ASCII code for NUL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken return the next token available for reading
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
