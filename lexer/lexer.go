package lexer

import (
	"fmt"
	"unicode"
)

type lexer struct {
	source string
	pos    int
	line   int
	Tokens []Token
}

func Tokenize(source string) []Token {
	lex := &lexer{
		source: source,
		pos:    0,
		line:   1,
		Tokens: make([]Token, 0),
	}

	for !lex.atEOF() {
		lex.scanToken()
	}

	return lex.Tokens
}

func (lex *lexer) scanToken() {
	ch := lex.peek()

	// skip whitespaces
	if unicode.IsSpace(rune(ch)) {
		lex.skipWhitespace()
		return
	}

	// comments
	if ch == '/' && lex.peek() == '/' {
		lex.skipComment()
		return
	}

	// string literals
	if ch == '"' {
		lex.scanString()
		return
	}

	if unicode.IsDigit(rune(ch)) {
		lex.scanNumber()
		return
	}

	// identifier and keywords
	if unicode.IsLetter(rune(ch)) || ch == '_' || ch == '-' {
		lex.scanIdentifier()
		return
	}

	switch ch {
	case '=', ':':
		lex.advance()
		lex.push(newUniqueToken(ASSIGN, "="))
		return
	case '[':
		lex.advance()
		lex.push(newUniqueToken(OPEN_BRACKET, "["))
		return
	case ']':
		lex.advance()
		lex.push(newUniqueToken(CLOSE_BRACKET, "]"))
		return
	}

	panic(fmt.Sprintf("lexer error: unexped character '%c' at position %d", ch, lex.pos))
}

func (lex *lexer) scanString() {
	start := lex.pos
	lex.advance()

	for !lex.atEOF() && lex.peek() != '"' {
		lex.advance()
	}

	if lex.atEOF() {
		panic("lexer error: unterminated string literal")
	}

	lex.advance()
	value := lex.source[start:lex.pos]
	lex.push(newUniqueToken(STRING, value))
}

func (lex *lexer) scanNumber() {
	start := lex.pos
	lex.advance()

	for !lex.atEOF() && unicode.IsDigit(rune(lex.peek())) {
		lex.advance()
	}

	if !lex.atEOF() && lex.peek() == '.' && unicode.IsDigit(rune(lex.peekNext())) {
		lex.advance()
		for !lex.atEOF() && unicode.IsDigit(rune(lex.peekNext())) {
			lex.advance()
		}
	}

	value := lex.source[start:lex.pos]
	lex.push(newUniqueToken(NUMBER, value))
}

func (lex *lexer) scanIdentifier() {
	start := lex.pos

	for !lex.atEOF() {
		ch := lex.peek()
		if unicode.IsLetter(rune(ch)) || unicode.IsDigit(rune(ch)) || ch == '_' || ch == '-' {
			lex.advance()
		} else {
			break
		}
	}

	value := lex.source[start:lex.pos]
	if kind, found := reservedWords[value]; found {
		lex.push(newUniqueToken(kind, value))
	} else {
		lex.push(newUniqueToken(IDENTIFIER, value))
	}
}

func (lex *lexer) skipWhitespace() {
	for !lex.atEOF() && unicode.IsSpace(rune(lex.peek())) {
		if lex.peek() == '\n' {
			lex.line++
		}
		lex.advance()
	}
}

func (lex *lexer) skipComment() {
	for !lex.atEOF() && lex.peek() != '\n' {
		lex.advance()
	}

	if !lex.atEOF() {
		lex.advance()
		lex.line++
	}
}

func (lex *lexer) peek() byte {
	if lex.atEOF() {
		return 0
	}

	return lex.source[lex.pos]
}

func (lex *lexer) peekNext() byte {
	if lex.pos+1 >= len(lex.source) {
		return 0
	}

	return lex.source[lex.pos+1]
}

func (lex *lexer) peekAhead(n int) byte {
	if lex.pos+n >= len(lex.source) {
		return 0
	}

	return lex.source[lex.pos+n]
}

func (lex *lexer) advance() {
	lex.pos++
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) atEOF() bool {
	return lex.pos >= len(lex.source)
}
