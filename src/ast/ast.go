package ast

import "github.com/tymbaca/gomonkey/src/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Stmts []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Stmts) > 0 {
		return p.Stmts[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetSatement struct {
	Tok   token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetSatement) TokenLiteral() string { return ls.Tok.Literal }
func (ls *LetSatement) statementNode()       {}

type Identifier struct {
	Tok   token.Token
	Value string // huh?
}

func (i *Identifier) TokenLiteral() string { return i.Tok.Literal }
func (i *Identifier) expressionNode()      {}
