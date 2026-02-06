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
	return &parser{
		tokens: tokens,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	_ = crerateParser(tokens)
	body := make([]ast.Stmt, 0)

	return ast.BlockStmt{
		Body: body,
	}
}
