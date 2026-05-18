package lexer

import (
	"interpreter/src/monkey/token"
)

const EOF_NUMBER = 0

type Lexer struct {
	input                         string
	currentPositionInsideTheInput int
	nextPositionToRead            int
	currentChar                   byte //same as input[currentPositionInsideTheInput]
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() *token.Token {
	var t *token.Token

	switch l.currentChar {
	case EOF_NUMBER: //NIL
		t = &token.Token{Type: token.EOF, Literal: ""}
	case ',':
		t = token.New(token.COMMA, l.currentChar)
	case ';':
		t = token.New(token.SEMICOLON, l.currentChar)
	case '(':
		t = token.New(token.LPAREN, l.currentChar)
	case ')':
		t = token.New(token.RPAREN, l.currentChar)
	case '{':
		t = token.New(token.LBRACE, l.currentChar)
	case '}':
		t = token.New(token.RBRACE, l.currentChar)
	case '=':
		t = token.New(token.ASSIGN, l.currentChar)
	case '+':
		t = token.New(token.PLUS, l.currentChar)
	default:
		t = l.processUnknowSymbol()
	}
	l.readChar()
	return t
}

func (l *Lexer) processUnknowSymbol() *token.Token {
	if l.isCurrentCharANumber() {
		return l.processNumber()
	} else if l.isCurrentCharALetter() {
		return l.processIdentifier()
	} else {
		return token.New(token.ILLEGAL, l.currentChar)
	}
}

func (l *Lexer) isCurrentCharALetter() bool {
	return 'a' <= l.currentChar && l.currentChar <= 'z' || 'A' <= l.currentChar && l.currentChar <= 'Z' || l.currentChar == '_'
}

func (l *Lexer) isCurrentCharANumber() bool {
	return l.currentChar >= '0' && l.currentChar <= '9'
}

func (l *Lexer) processNumber() *token.Token {
	position := l.currentPositionInsideTheInput
	for l.isCurrentCharANumber() {
		l.readChar()
	}
	return token.NewTokenWithLiteral(token.INT, l.input[position:l.currentPositionInsideTheInput])
}

func (l *Lexer) processIdentifier() *token.Token {
	position := l.currentPositionInsideTheInput
	for l.isCurrentCharALetter() || l.isCurrentCharANumber() {
		l.readChar()
	}
	return token.NewTokenWithLiteral(token.INT, l.input[position:l.currentPositionInsideTheInput])
}

func (l *Lexer) readChar() {
	l.updateCurrentChar()
	l.updatePositions()
}

func (l *Lexer) updateCurrentChar() {
	if l.nextPositionToRead >= len(l.input) {
		l.currentChar = EOF_NUMBER
	} else {
		l.currentChar = l.input[l.nextPositionToRead]
	}
}

func (l *Lexer) updatePositions() {
	l.currentPositionInsideTheInput = l.nextPositionToRead
	l.nextPositionToRead += 1
}
