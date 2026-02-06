package parser

import (
	"client_http/ast"
	"client_http/lexer"
)

type bindingPower int

const (
	default_bp bindingPower = iota
	assignment
	primary
	unary
	member
)

type (
	stmtHandler func(p *parser) ast.Stmt
	nudHandler  func(p *parser) ast.Expr
	ledHandler  func(p *parser, left ast.Expr, bp bindingPower) ast.Expr

	stmtLookup map[lexer.TokenKind]stmtHandler
	nudLookup  map[lexer.TokenKind]nudHandler
	ledLookup  map[lexer.TokenKind]ledHandler
	bpLookup   map[lexer.TokenKind]bindingPower
)

var (
	bpLu   = bpLookup{}
	nudLu  = nudLookup{}
	ledLu  = ledLookup{}
	stmtLu = stmtLookup{}
)

func led(kind lexer.TokenKind, bp bindingPower, ledFn ledHandler) {
	bpLu[kind] = bp
	ledLu[kind] = ledFn
}

func nud(kind lexer.TokenKind, nudFn nudHandler) {
	nudLu[kind] = nudFn
}

func stmt(kind lexer.TokenKind, stmtFn stmtHandler) {
	bpLu[kind] = default_bp
	stmtLu[kind] = stmtFn
}

func createTokenLookups() {

	led(lexer.ASSIGN, assignment, parseAssignmentExpr)
	led(lexer.OPEN_BRACKET, member, parseMemberExpr)

	nud(lexer.IDENTIFIER, parsePrimaryExpr)
	nud(lexer.STRING, parsePrimaryExpr)
	nud(lexer.NUMBER, parsePrimaryExpr)

	stmt(lexer.METHOD, parseMethodStmt)
	stmt(lexer.OPEN_BRACKET, parseBlockStmt)
}
