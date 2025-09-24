package strategies

import (
	"lang/lexer"
)

type SpaceStrategy struct{}

func (ss *SpaceStrategy) Try(input string, pos int) (bool, int, *lexer.Token) {
	char := string(input[pos])
	if char == " " || char == "\n" {
		return true, pos, &lexer.Token{Literal: "", TokenType: lexer.SPACE}
	}

	return false, pos, nil
}
