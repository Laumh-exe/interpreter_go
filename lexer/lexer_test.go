package lexer

import (
	"github.com/Laumh-exe/interpreter_go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input :=
		`let five = 5; 
		let ten = 10; 
		let add = fn(x, y) {
		x + y 
		};`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{expectedType: token.IDENT, expectedLiteral: "five"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "5"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENT, expectedLiteral: "ten"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "10"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENT, expectedLiteral: "add"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.FUNCTION, expectedLiteral: "fn"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENT, expectedLiteral: "x"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENT, expectedLiteral: "y"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.LBRACE, expectedLiteral: "{"},
		{expectedType: token.IDENT, expectedLiteral: "x"},
		{expectedType: token.PLUS, expectedLiteral: "+"},
		{expectedType: token.IDENT, expectedLiteral: "y"},
		{expectedType: token.RBRACE, expectedLiteral: "}"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - toketype.wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - toketype.wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
