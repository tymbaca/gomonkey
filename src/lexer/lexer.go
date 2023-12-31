package lexer

import "github.com/tymbaca/gomonkey/src/token"

type Lexer struct {
	input   string
	pos     int
	readPos int
	char    byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case '[':
		tok = newToken(token.LBRACK, l.char)
	case ']':
		tok = newToken(token.RBRACK, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}

	l.readChar()
	return tok
}

func newToken(typ token.TokenType, lit byte) token.Token {
	return token.Token{Type: typ, Literal: string(lit)}
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos++
}
