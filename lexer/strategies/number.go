package strategies

import (
	"lang/lexer"
	"unicode"
)

type NumberStrategy struct{}

func (ns *NumberStrategy) Try(input string, pos int) (bool, int, *lexer.Token) { // TODO: contemplar també els decimals que contenen punts
	ok := unicode.IsDigit(rune(input[pos]))
	if !ok {
		return false, pos, nil
	}

	size := len(input)
	count := pos + 1
	for count < size {
		ok = unicode.IsDigit(rune(input[count]))
		if !ok {
			break
		}
		count++
	}

	if count >= size {
		return true, count, &lexer.Token{Literal: input[pos:count], TokenType: lexer.NUMBER}
	}
	if input[count] == '.' {
		count++
		for count < size {
			ok = unicode.IsDigit(rune(input[count]))
			if !ok {
				break
			}
			count++
		}
	}

	return true, count - 1, &lexer.Token{Literal: input[pos:count], TokenType: lexer.NUMBER}
}
