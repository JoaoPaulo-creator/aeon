package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	IDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET

	// reserved words
	METHOD
	BODY
	HEADERS

	// =
	ASSIGN
)

var reservedWords map[string]TokenKind = map[string]TokenKind{
	"method":  METHOD,
	"body":    BODY,
	"headers": HEADERS,
	"assign":  ASSIGN,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case METHOD:
		return "method"
	case BODY:
		return "body"
	case HEADERS:
		return "headers"
	case IDENTIFIER:
		return "identifier"
	case ASSIGN:
		return "assign"
	default:
		return fmt.Sprintf("unknown token: %d", kind)
	}
}

func newUniqueToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}
