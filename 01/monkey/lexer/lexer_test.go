package lexer

import (
	"testing"

	"github.com/francojposa/interpreter-in-go/01/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype incorrect. expected: %q, got: %q",
				i, tt.expectedType, token.Type,
			)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal incorrect. expected: %q, got: %q",
				i, tt.expectedLiteral, token.Literal,
			)
		}
	}
}
