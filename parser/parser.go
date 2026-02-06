package parser

import (
	"client_http/ast"
	"client_http/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func crerateParser(tokens []lexer.Token) *parser {
	createTokenLookups()

	return &parser{
		tokens: tokens,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	p := crerateParser(tokens)
	body := make([]ast.Stmt, 0)

	for p.hasTokens() {
		body = append(body, parseStmt(p))
	}

	return ast.BlockStmt{
		Body: body,
	}
}
