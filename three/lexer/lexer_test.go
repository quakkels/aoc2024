package lexer

import (
	"testing"
	"three/token"
)

func TestNextToken(t *testing.T) {
	input := "1(,)"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "1"},
		{token.LPAREN, "("},
		{token.COMMA, ","},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatal("expected:", tt.expectedType, "actual:", tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatal("expect:", tt.expectedLiteral, "actual:", tok.Literal)
		}
	}
}
