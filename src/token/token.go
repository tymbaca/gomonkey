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

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
