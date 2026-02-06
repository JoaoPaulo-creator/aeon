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
	BODY
	HEADERS
	POST
	GET
	PUT

	// = or :
	ASSIGN
)

var reservedWords map[string]TokenKind = map[string]TokenKind{
	"method":  METHOD,
	"body":    BODY,
	"headers": HEADERS,
	"post":    POST,
	"get":     GET,
	"put":     PUT,
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
	case STRING:
		return "string"
	case NUMBER:
		return "number"
	case IDENTIFIER:
		return "identifier"
	case ASSIGN:
		return "assign"
	case POST:
		return "post"
	case GET:
		return "get"
	case PUT:
		return "put"
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
