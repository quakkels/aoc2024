package lexer

import (
	"testing"
	"three/token"
)

func TestNextToken(t *testing.T) {
	input := " mul(123,3)"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.MUL, "mul"},
		{token.LPAREN, "("},
		{token.INT, "123"},
		{token.COMMA, ","},
		{token.INT, "3"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()
		t.Log("Literal:", tok.Literal)
		if tok.Type != tt.expectedType {
			t.Fatal("expected:", tt.expectedType, "actual:", tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatal("expect:", tt.expectedLiteral, "actual:", tok.Literal)
		}
	}
}
