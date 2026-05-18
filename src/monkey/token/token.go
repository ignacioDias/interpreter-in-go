package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // names for vars, func,...
	INT   = "INT"   // 1343456
	TRUE  = "TRUE"
	FALSE = "FALSE"
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	BANG     = "!"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	IF       = "IF"
	RETURN   = "RETURN"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"var": VAR,
}

func New(tokenType TokenType, char byte) *Token {
	return &Token{Type: tokenType, Literal: string(char)}
}

func NewTokenWithLiteral(tokenType TokenType, str string) *Token {
	return &Token{Type: tokenType, Literal: str}
}

func GenerateTokenForWord(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
