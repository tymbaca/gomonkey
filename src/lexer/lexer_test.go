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
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong. expected=%q, got=%q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
