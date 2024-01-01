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
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
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
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "five"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},

				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "ten"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "10"},
				{Type: token.SEMICOLON, Literal: ";"},

				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "add"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.FUNCTION, Literal: "fn"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},

				{Type: token.IDENT, Literal: "x"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.SEMICOLON, Literal: ";"},

				{Type: token.RBRACE, Literal: "}"},
				{Type: token.SEMICOLON, Literal: ";"},

				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "result"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.IDENT, Literal: "add"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "five"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.IDENT, Literal: "ten"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.SEMICOLON, Literal: ";"},

				{Type: token.EOF, Literal: ""},
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
				{Type: token.BANG, Literal: token.BANG},
				{Type: token.MINUS, Literal: token.MINUS},
				{Type: token.SLASH, Literal: token.SLASH},
				{Type: token.ASTERISK, Literal: token.ASTERISK},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: token.SEMICOLON},

				{Type: token.INT, Literal: "5"},
				{Type: token.LT, Literal: token.LT},
				{Type: token.INT, Literal: "10"},
				{Type: token.GT, Literal: token.GT},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: token.SEMICOLON},

				{Type: token.IDENT, Literal: "age"},
				{Type: token.LE, Literal: token.LE},
				{Type: token.INT, Literal: "16"},
				{Type: token.SEMICOLON, Literal: token.SEMICOLON},

				{Type: token.IDENT, Literal: "age"},
				{Type: token.GE, Literal: token.GE},
				{Type: token.INT, Literal: "12"},
				{Type: token.SEMICOLON, Literal: token.SEMICOLON},

				{Type: token.IDENT, Literal: "age"},
				{Type: token.NOT_EQ, Literal: token.NOT_EQ},
				{Type: token.INT, Literal: "13"},
				{Type: token.SEMICOLON, Literal: token.SEMICOLON},

				{Type: token.IDENT, Literal: "age"},
				{Type: token.EQ, Literal: token.EQ},
				{Type: token.INT, Literal: "15"},
				{Type: token.SEMICOLON, Literal: token.SEMICOLON},
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
				{Type: token.IF, Literal: "if"},

				{Type: token.LPAREN, Literal: token.LPAREN},
				{Type: token.IDENT, Literal: "age"},
				{Type: token.LT, Literal: token.LT},
				{Type: token.INT, Literal: "10"},
				{Type: token.RPAREN, Literal: token.RPAREN},

				{Type: token.LBRACE, Literal: token.LBRACE},

				// 2 line
				{Type: token.RETURN, Literal: "return"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.SEMICOLON, Literal: token.SEMICOLON},

				// 3 line
				{Type: token.RBRACE, Literal: token.RBRACE},
				{Type: token.ELSE, Literal: "else"},
				{Type: token.LBRACE, Literal: token.LBRACE},

				// 4 line
				{Type: token.RETURN, Literal: "return"},
				{Type: token.FALSE, Literal: "false"},
				{Type: token.SEMICOLON, Literal: token.SEMICOLON},

				// 5 line
				{Type: token.RBRACE, Literal: token.RBRACE},
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
