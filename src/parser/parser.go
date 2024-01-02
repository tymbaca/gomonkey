package parser

import (
	"github.com/tymbaca/gomonkey/src/ast"
	"github.com/tymbaca/gomonkey/src/lexer"
	"github.com/tymbaca/gomonkey/src/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// need to reat 2 times to get correct start state (cur will look at 1st and peek will look at 2nd)
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
