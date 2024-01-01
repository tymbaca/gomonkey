package token

const (
	// Special
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Vars shit (identifiers and literals)
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	PLUS   = "+"
	ASSIGN = "="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// Brackets
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var (
	keywords = map[string]TokenType{
		"fn":  FUNCTION,
		"let": LET,
		"var": LET,
	}
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func GetIdentOrKeyword(word string) TokenType {
	if keyword, ok := keywords[word]; ok {
		return keyword
	}
	return IDENT
}
