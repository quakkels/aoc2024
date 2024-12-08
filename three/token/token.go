package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT"
	COMMA   = ","
	LPAREN  = "("
	RPAREN  = ")"

	// keyword types
	MUL = "MUL"
)

var keywords = map[string]TokenType {
	"mul": MUL,
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func LookupIdentType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return ILLEGAL
}

