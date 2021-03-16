package lexer

type Lexer struct {
	ch               byte
	input            string
	position, offset int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.offset >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.offset]
	}

	l.position = l.offset
	l.offset++
}
