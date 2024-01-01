package lexer

import (
	"reflect"
	"testing"

	"github.com/tymbaca/gomonkey/src/token"
)

func TestNew(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *Lexer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestLexer_NextToken(t *testing.T) {
// 	type fields struct {
// 		input string
// 		pos   int
// 	}
// 	tests := []struct {
// 		input string
// 		want  token.Token
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			input: "=",
// 			want:  token.Token{token.ASSIGN, "="},
// 		},
// 		{
// 			input: "+",
// 			want:  token.Token{token.PLUS, "+"},
// 		},
// 		{
// 			input: "(",
// 			want:  token.Token{token.LPAREN, "("},
// 		},
// 		{
// 			input: ")",
// 			want:  token.Token{token.RPAREN, ")"},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &Lexer{
// 				input: tt.fields.input,
// 				pos:   tt.fields.pos,
// 			}
// 			if got := l.NextToken(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Lexer.NextToken() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestLexer_NextToken(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		expToks []token.Token
	}{
		{
			name:  "first symbols",
			input: `=+(){},;`,
			expToks: []token.Token{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name: "simple program",
			input: `let five = 5;
let ten = 10;

let add = fn(x, y) {
        x + y;
};
let result = add(five, ten);`,
			expToks: []token.Token{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},

				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},

				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},

				{token.EOF, ""},
			},
		},
		{
			name: "additional symbols and comparison",
			input: `
                                !-/*5;
                                5 < 10 > 5;
                                age <= 16;
                                age >= 12;
                                age != 13;
                                age == 15;
                        `,
			expToks: []token.Token{
				{token.BANG, token.BANG},
				{token.MINUS, token.MINUS},
				{token.SLASH, token.SLASH},
				{token.ASTERISK, token.ASTERISK},
				{token.INT, "5"},
				{token.SEMICOLON, token.SEMICOLON},

				{token.INT, "5"},
				{token.LT, token.LT},
				{token.INT, "10"},
				{token.GT, token.GT},
				{token.INT, "5"},
				{token.SEMICOLON, token.SEMICOLON},

				{token.IDENT, "age"},
				{token.LE, token.LE},
				{token.INT, "16"},
				{token.SEMICOLON, token.SEMICOLON},

				{token.IDENT, "age"},
				{token.GE, token.GE},
				{token.INT, "12"},
				{token.SEMICOLON, token.SEMICOLON},

				{token.IDENT, "age"},
				{token.NOT_EQ, token.NOT_EQ},
				{token.INT, "13"},
				{token.SEMICOLON, token.SEMICOLON},

				{token.IDENT, "age"},
				{token.EQ, token.EQ},
				{token.INT, "15"},
				{token.SEMICOLON, token.SEMICOLON},
			},
		},
		{
			name: "new keywords",
			input: `
                        if (age < 10) {
                                return true;
                        } else {
                                return false;
                        }
                        `,
			expToks: []token.Token{
				// 1 line
				{token.IF, "if"},

				{token.LPAREN, token.LPAREN},
				{token.IDENT, "age"},
				{token.LT, token.LT},
				{token.INT, "10"},
				{token.RPAREN, token.RPAREN},

				{token.LBRACE, token.LBRACE},

				// 2 line
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, token.SEMICOLON},

				// 3 line
				{token.RBRACE, token.RBRACE},
				{token.ELSE, "else"},
				{token.LBRACE, token.LBRACE},

				// 4 line
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, token.SEMICOLON},

				// 5 line
				{token.RBRACE, token.RBRACE},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)
			for ii, expTok := range tt.expToks {
				tok := l.NextToken()
				if tok.Type != expTok.Type {
					t.Fatalf(
						"tests[%d] - tokentype wrong. expected=%q, got=%q",
						ii,
						expTok.Type,
						tok.Type,
					)
				}
				if tok.Literal != expTok.Literal {
					t.Fatalf(
						"tests[%d] - literal wrong. expected=%q, got=%q",
						ii,
						expTok.Literal,
						tok.Literal,
					)
				}
				// t.Logf("tests[%d] - token %q parsed correctly", ii, tok)
			}
			// t.Logf(
			// 	"test[%d] - passed with input: %s",
			// 	i,
			// 	tt.input,
			// )
		})
	}
}

func TestLexer_readChar(t *testing.T) {
	type fields struct {
		input   string
		pos     int
		readPos int
		char    byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:   tt.fields.input,
				pos:     tt.fields.pos,
				readPos: tt.fields.readPos,
				char:    tt.fields.char,
			}
			l.readChar()
		})
	}
}

func TestLexer_readIdentifier(t *testing.T) {
	type fields struct {
		input   string
		pos     int
		readPos int
		char    byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:   tt.fields.input,
				pos:     tt.fields.pos,
				readPos: tt.fields.readPos,
				char:    tt.fields.char,
			}
			if got := l.readLetters(); got != tt.want {
				t.Errorf("Lexer.readIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_skipWhitespaces(t *testing.T) {
	type fields struct {
		input   string
		pos     int
		readPos int
		char    byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:   tt.fields.input,
				pos:     tt.fields.pos,
				readPos: tt.fields.readPos,
				char:    tt.fields.char,
			}
			l.skipWhitespaces()
		})
	}
}

func Test_newToken(t *testing.T) {
	type args struct {
		typ token.TokenType
		lit byte
	}
	tests := []struct {
		name string
		args args
		want token.Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newToken(tt.args.typ, tt.args.lit); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("newToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLetter(t *testing.T) {
	tests := []struct {
		charset []byte
		want    bool
	}{
		// TODO: Add test cases.
		{
			[]byte(`qwertyuiopasdfghjklzxcvbnm_QWERTYUIOPASDFGHJKLZXCVBNM`),
			true,
		},
		{
			[]byte("`\"'[]{}()<>;:,./?!\\|-=+#%^*~1234567890"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.charset), func(t *testing.T) {
			for _, char := range tt.charset {
				if got := isLetter(char); got != tt.want {
					t.Errorf("isLetter() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_isDigit(t *testing.T) {
	tests := []struct {
		charset []byte
		want    bool
	}{
		// TODO: Add test cases.
		{
			[]byte(`1234567890`),
			true,
		},
		{
			[]byte(
				"qwertyuiopasdfghjklzxcvbnm_QWERTYUIOPASDFGHJKLZXCVBNM`\"'[]{}()<>;:,./?!\\|-=+#%^*~",
			),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.charset), func(t *testing.T) {
			for _, char := range tt.charset {
				if got := isDigit(char); got != tt.want {
					t.Errorf("isDigit() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
