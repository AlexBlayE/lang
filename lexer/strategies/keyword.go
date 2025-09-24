package strategies

import (
	"lang/lexer"
	"unicode"
)

type KeywordStrategy struct{} // TODO: fer un keyword dispatcher

func (ks *KeywordStrategy) Try(input string, pos int) (bool, int, *lexer.Token) {
	char := string(input[pos])

	elem, ok := lexer.Keywords[char]
	if ok {
		return true, pos, &lexer.Token{Literal: char, TokenType: elem}
	}

	size := len(input)
	count := pos + 1
	for count < size {
		if !unicode.IsLetter(rune(input[count])) {
			elem, ok = lexer.Keywords[input[pos:count]]
			if ok {
				return true, count - 1, &lexer.Token{Literal: input[pos:count], TokenType: elem}
			}

			return true, count - 1, &lexer.Token{Literal: input[pos:count], TokenType: lexer.LITERAL}
		}
		count++
	}

	return false, pos, nil
}
