package ast

import "interpreter/src/monkey/token"

type Node interface {
	TokenLiteral() string //for debugging and testing
}

type StatementNode interface {
	Node
}

type ExpressionNode interface {
	Node
}

type Program struct {
	Statements []StatementNode
}

type VarStatement struct {
	Token      token.Token
	Name       string
	Expression *ExpressionNode
}

type Identifier struct {
	Token token.Token
	Value string
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return "nil"
}

func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
