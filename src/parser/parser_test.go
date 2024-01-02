package parser

import (
	"reflect"
	"testing"

	"github.com/tymbaca/gomonkey/src/ast"
	"github.com/tymbaca/gomonkey/src/lexer"
	"github.com/tymbaca/gomonkey/src/token"
)

func TestLetStatement(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expIdents []string
	}{
		// TODO: Add test cases.
		{
			name: "first",
			input: `
                        let name = 6;
                        let age = 18;
                        let foot = 300;`,
			expIdents: []string{"name", "age", "foot"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)

			program := p.ParseProgram()

			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			if len(program.Stmts) != len(tt.expIdents) {
				t.Fatalf(
					"incorrect statements amount, need: %d, got: %d",
					len(tt.expIdents),
					len(program.Stmts),
				)
			}

			for i := range program.Stmts {
				stmt := program.Stmts[i]

				if stmt.TokenLiteral() != "let" {
					t.Errorf(
						"statement literal doesn't start with 'let', got: %s",
						stmt.TokenLiteral(),
					)
				}

				letStmt, ok := stmt.(*ast.LetSatement)
				if !ok {
					t.Errorf(
						"statement not of type 'let', got type: %T",
						stmt,
					)
				}

				if letStmt.Name.Value != tt.expIdents[i] {
					t.Errorf(
						"wrong statement name epected: %s, got type: %s",
						tt.expIdents[i],
						letStmt.Name.Value,
					)
				}
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		l *lexer.Lexer
	}
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_nextToken(t *testing.T) {
	type fields struct {
		l         *lexer.Lexer
		curToken  token.Token
		peekToken token.Token
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				l:         tt.fields.l,
				curToken:  tt.fields.curToken,
				peekToken: tt.fields.peekToken,
			}
			p.nextToken()
		})
	}
}

func TestParser_ParseProgram(t *testing.T) {
	type fields struct {
		l         *lexer.Lexer
		curToken  token.Token
		peekToken token.Token
	}
	tests := []struct {
		name   string
		fields fields
		want   *ast.Program
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				l:         tt.fields.l,
				curToken:  tt.fields.curToken,
				peekToken: tt.fields.peekToken,
			}
			if got := p.ParseProgram(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.ParseProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}
