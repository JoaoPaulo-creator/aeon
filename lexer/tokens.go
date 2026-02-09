package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	IDENTIFIER
	STRING
	NUMBER

	OPEN_BRACKET
	CLOSE_BRACKET

	// reserved words
	METHOD
	ENDPOINT
	BODY
	HEADERS

	// = or :
	ASSIGN
)

var reservedWords map[string]TokenKind = map[string]TokenKind{
	"method":   METHOD,
	"endpoint": ENDPOINT,
	"body":     BODY,
	"headers":  HEADERS,
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
	case ENDPOINT:
		return "enpoint"
	case BODY:
		return "body"
	case HEADERS:
		return "headers"
	case STRING:
		return "string"
	case NUMBER:
		return "number"
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
