package lexer

type Token struct {
	Literal   string
	TokenType TokenType
}

type Strategy interface {
	Try(str string, pos int) (bool, int, *Token)
}

type Lexer struct {
	strategies []Strategy
	content    string
	channel    chan *Token
	pos        int
}

func NewLexer(input string, strategies []Strategy) *Lexer {
	return &Lexer{
		strategies: strategies,
		content:    input,
		pos:        0,
		channel:    make(chan *Token),
	}
}

func (l *Lexer) NextToken() <-chan *Token {
	go l.startLexing()
	return l.channel
}

func (l *Lexer) startLexing() {
	// defer close(l.channel)
	size := len(l.content)

	for l.pos < size {
		for _, strategy := range l.strategies {
			ok, newPos, tok := strategy.Try(l.content, l.pos)
			if ok {
				l.channel <- tok
				l.pos = newPos
				break
			}
		}
		l.pos++
	}

	l.channel <- &Token{Literal: "", TokenType: EOF}
}
