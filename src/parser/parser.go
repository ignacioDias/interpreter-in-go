package parser

import (
	"interpreter/src/ast"
	"interpreter/src/monkey/lexer"
	"interpreter/src/monkey/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	nextToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}

	p.updateCurrentTokens()
	p.updateCurrentTokens()
	return p
}

func (p *Parser) updateCurrentTokens() {
	p.currentToken = p.nextToken
	p.nextToken = *p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
