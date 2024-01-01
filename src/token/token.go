package token

const (
	// Special
	BANG    = "!"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Vars shit (identifiers and literals)
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	ASSIGN   = "="
	EQ       = "=="
	NOT_EQ   = "!="
	LT       = "<"
	GT       = ">"
	LE       = "<="
	GE       = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// Brackets
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// LBRACK = "["
	// RBRACK = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	IF     = "IF"
	ELSE   = "ELSE"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	RETURN = "RETURN"
)

var (
	keywords = map[string]TokenType{
		"func": FUNCTION,
		"fn":   FUNCTION,
		"let":  LET,
		"var":  LET,

		"if":     IF,
		"else":   ELSE,
		"true":   TRUE,
		"false":  FALSE,
		"return": RETURN,
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
