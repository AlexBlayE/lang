package lexer_test

import (
	"lang/lexer"
	"lang/lexer/strategies"
	"testing"
)

func TestLexer(t *testing.T) {
	statements := []string{
		"let num = 5 * 2",
		"let num = duplicate(5) + 10",
		"fn do() { return 5 }",
		"for i -> 5 {}",
	}

	responses := [][]lexer.TokenType{
		{lexer.LET, lexer.SPACE, lexer.LITERAL, lexer.SPACE, lexer.EQUALS, lexer.SPACE, lexer.NUMBER, lexer.SPACE, lexer.MUL, lexer.SPACE, lexer.NUMBER},
		{lexer.LET, lexer.SPACE, lexer.LITERAL, lexer.SPACE, lexer.EQUALS, lexer.SPACE, lexer.LITERAL, lexer.LPAREN, lexer.NUMBER, lexer.RPAREN, lexer.SPACE, lexer.ADD, lexer.SPACE, lexer.NUMBER},
		{lexer.FUNC, lexer.SPACE, lexer.LITERAL, lexer.LPAREN, lexer.RPAREN, lexer.SPACE, lexer.LBRACE, lexer.SPACE, lexer.RETURN, lexer.SPACE, lexer.NUMBER, lexer.SPACE, lexer.RBRACE},
		{lexer.FOR, lexer.SPACE, lexer.LITERAL, lexer.SPACE, lexer.MIN, lexer.GTHAN, lexer.SPACE, lexer.NUMBER, lexer.SPACE, lexer.LBRACE, lexer.RBRACE},
	}

	for i, stm := range statements {
		lex := lexer.NewLexer(stm, []lexer.Strategy{
			&strategies.SpaceStrategy{},
			&strategies.NumberStrategy{},
			&strategies.KeywordStrategy{},
		})

		count := 0
		for tok := range lex.NextToken() {
			if tok.TokenType == lexer.EOF {
				break
			}

			res := responses[i][count]

			if tok.TokenType != res {
				t.Error("\ntok.TokenType: "+tok.TokenType+"\ntok.Literal: "+lexer.TokenType(tok.Literal)+"\nRes: ", res)
			}

			count++
		}
	}
}
