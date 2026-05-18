package lexer

import (
	"interpreter/src/monkey/token"
)

//ASCII

const REPRESENTATION_FOR_NUMBER_ZERO = 48
const REPRESENTATION_FOR_NUMBER_NINE = 57

const REPRESENTATION_FOR_CAPITAL_A = 65
const REPRESENTATION_FOR_CAPITAL_Z = 90

const REPRESENTATION_FOR_LOWER_A = 97
const REPRESENTATION_FOR_LOWER_Z = 122

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
	case 0: //NIL
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

func (l *Lexer) readChar() {
	l.updateCurrentChar()
	l.updatePositions()
}

func (l *Lexer) updateCurrentChar() {
	if l.nextPositionToRead >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.nextPositionToRead]
	}
}

func (l *Lexer) updatePositions() {
	l.currentPositionInsideTheInput = l.nextPositionToRead
	l.nextPositionToRead += 1
}

func (l *Lexer) processUnknowSymbol() *token.Token {
	if l.isCurrentCharANumber() {
		return l.processNumber()
	} else if l.isCurrentCharALetter() {
		return l.processWord()
	} else {
		return token.New(token.ILLEGAL, l.currentChar)
	}
}

func (l *Lexer) isCurrentCharALetter() bool {
	var isLowercaseLetter = l.currentChar >= REPRESENTATION_FOR_LOWER_A || l.currentChar <= REPRESENTATION_FOR_LOWER_Z
	var isUppercaseLetter = l.currentChar >= REPRESENTATION_FOR_CAPITAL_A || l.currentChar <= REPRESENTATION_FOR_CAPITAL_Z
	return isLowercaseLetter || isUppercaseLetter
}

func (l *Lexer) isCurrentCharANumber() bool {
	return l.currentChar >= REPRESENTATION_FOR_NUMBER_ZERO && l.currentChar <= REPRESENTATION_FOR_NUMBER_NINE
}

func (l *Lexer) processNumber() *token.Token {

}
