package lexer

type TokenType string

const (
	LITERAL TokenType = "literal"
	NUMBER  TokenType = "number"

	TRUE  TokenType = "true"
	FALSE TokenType = "false"

	RETURN TokenType = "return"

	LET  TokenType = "let"
	FUNC TokenType = "fn"
	IF   TokenType = "if"
	FOR  TokenType = "for"

	LBRACE TokenType = "{"
	RBRACE TokenType = "}"
	LPAREN TokenType = "("
	RPAREN TokenType = ")"

	EQUALS TokenType = "="
	GTHAN  TokenType = ">"
	LTHAN  TokenType = "<"

	ADD TokenType = "+"
	MIN TokenType = "-"
	DIV TokenType = "/"
	MUL TokenType = "*"

	STRDELIM TokenType = "\""

	SPACE TokenType = "space"

	EOF TokenType = "EOF"
)
