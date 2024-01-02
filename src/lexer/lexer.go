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

	l.skipWhitespaces()

	switch {
	case l.char == '(':
		tok = newToken(token.LPAREN, l.char)
	case l.char == ')':
		tok = newToken(token.RPAREN, l.char)
	case l.char == '{':
		tok = newToken(token.LBRACE, l.char)
	case l.char == '}':
		tok = newToken(token.RBRACE, l.char)
	// case l.char == '[':
	// 	tok = newToken(token.LBRACK, l.char)
	// case l.char == ']':
	// 	tok = newToken(token.RBRACK, l.char)

	case l.char == ',':
		tok = newToken(token.COMMA, l.char)
	case l.char == ';':
		tok = newToken(token.SEMICOLON, l.char)

	case l.char == '+':
		tok = newToken(token.PLUS, l.char)
	case l.char == '-':
		tok = newToken(token.MINUS, l.char)
	case l.char == '*':
		tok = newToken(token.ASTERISK, l.char)
	case l.char == '/':
		tok = newToken(token.SLASH, l.char)

	case l.char == '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: l.input[l.pos-1 : l.pos+1],
			}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case l.char == '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: l.input[l.pos-1 : l.pos+1],
			}
		} else {
			// NOTE: is it legal?
			tok = newToken(token.BANG, l.char)
		}
	case l.char == '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{
				Type:    token.LE,
				Literal: l.input[l.pos-1 : l.pos+1],
			}
		} else {
			tok = newToken(token.LT, l.char)
		}
	case l.char == '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{
				Type:    token.GE,
				Literal: l.input[l.pos-1 : l.pos+1],
			}
		} else {
			tok = newToken(token.GT, l.char)
		}

	case l.char == 0:
		tok.Type = token.EOF
		tok.Literal = ""
	case isLetter(l.char):
		tok.Literal = l.readLetters()
		tok.Type = token.GetIdentOrKeyword(tok.Literal)
		// needed because l.readLetters() already done readChar()
		// and we don't need to do it below
		return tok
	case isDigit(l.char):
		tok.Literal = l.readDigits()
		// WARN: handle "523someshit"
		// here we can check if nextChar eather " " or letter
		// also we can add support for "1_000_000" int literals
		tok.Type = token.INT
		return tok
	default:
		tok = newToken(token.ILLEGAL, l.char)
	}

	l.readChar()
	return tok
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

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Lexer) readLetters() string {
	startPos := l.pos

	for isLetter(l.char) {
		l.readChar()
	}

	// At this momet l.pos will look at non-letter, so we not include it
	return l.input[startPos:l.pos]
}

func (l *Lexer) readDigits() string {
	startPos := l.pos

	for isDigit(l.char) {
		l.readChar()
	}

	// At this momet l.pos will look at non-letter, so we not include it
	return l.input[startPos:l.pos]
}

func (l *Lexer) skipWhitespaces() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
	return
}

func newToken(typ token.TokenType, lit byte) token.Token {
	return token.Token{Type: typ, Literal: string(lit)}
}

// Maybe move this to NextToken() switch case
func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' ||
		char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
