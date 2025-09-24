package lexer

var Keywords map[string]TokenType = map[string]TokenType{
	"fn":     FUNC,
	"let":    LET,
	"=":      EQUALS,
	"+":      ADD,
	"-":      MIN,
	"*":      MUL,
	"/":      DIV,
	"if":     IF,
	"for":    FOR,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"(":      LPAREN,
	")":      RPAREN,
	"{":      LBRACE,
	"}":      RBRACE,
	"<":      LTHAN,
	">":      GTHAN,
}
