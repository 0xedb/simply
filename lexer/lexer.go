package lexer

import (
	"log"
	"regexp"

	"github.com/thebashshell/simply/token"
)

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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case 0:
		tok = token.Token{Kind: token.EOF}
	case '=':
		tok = makeToken(token.ASSIGN, l.ch)
	case '+':
		tok = makeToken(token.PLUS, l.ch)
	case ',':
		tok = makeToken(token.COMMA, l.ch)
	case ';':
		tok = makeToken(token.SEMICOLON, l.ch)
	case '(':
		tok = makeToken(token.LPAREN, l.ch)
	case ')':
		tok = makeToken(token.RPAREN, l.ch)
	case '{':
		tok = makeToken(token.LBRACE, l.ch)
	case '}':
		tok = makeToken(token.RBRACE, l.ch)

	default:
		if isDigit(l.ch) {
			tok.Value = l.readNumber()
			tok.Kind = token.INT
			return tok
		} else if isLetter(l.ch) {
			tok.Value = l.getIdentifier()
			tok.Kind = token.LookUpIdentifier(tok.Value)
			return tok
		} else {
			tok = makeToken(token.ILLEGAL, l.ch)
			return tok
		}

	}

	l.readChar()
	return tok
}

func makeToken(tok string, ch byte) token.Token {
	return token.Token{Kind: tok, Value: string(ch)}
}

func (l *Lexer) getIdentifier() string {
	pos := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func isLetter(ch byte) bool {
	ok, err := regexp.MatchString(`[a-zA-Z_]`, string(ch))

	if err != nil {
		log.Fatal(err)
	}

	return ok
}

func isDigit(ch byte) bool {
	ok, err := regexp.MatchString(`[0-9]`, string(ch))

	if err != nil {
		log.Fatal(err)
	}

	return ok
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
