package parser

import (
	"client_http/ast"
	"client_http/lexer"
)

func parseStmt(p *parser) ast.Stmt {
	stmtFn, exists := stmtLu[p.currentTokenKind()]
	if exists {
		return stmtFn(p)
	}

	expression := parseExpr(p, default_bp)
	return ast.ExpressionStmt{
		Expression: expression,
	}
}

func parseMethodStmt(p *parser) ast.Stmt {
	p.expect(lexer.METHOD)
	p.expect(lexer.ASSIGN)

	val := parseExpr(p, default_bp)
	return ast.MethodStmt{
		Value: val,
	}
}

func parseEndpointStmt(p *parser) ast.Stmt {
	p.expect(lexer.ENDPOINT)
	p.expect(lexer.ASSIGN)

	val := parseExpr(p, default_bp)
	return ast.EndpointStmt{
		Value: val,
	}
}

func parseHeadersStmt(p *parser) ast.Stmt {
	p.expect(lexer.HEADERS)
	if p.currentTokenKind() == lexer.ASSIGN {
		p.advance()

	}

	var properties = map[string]string{}
	p.expect(lexer.OPEN_BRACKET)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_BRACKET {
		var identName string

		identName = p.expect(lexer.IDENTIFIER).Value
		p.expectError(lexer.ASSIGN, "expected to find colon following property name")

		switch p.currentTokenKind() {
		case lexer.IDENTIFIER, lexer.STRING, lexer.NUMBER:
			properties[identName] = p.advance().Value
		default:
			panic("expected header valur after ':'")
		}

		// properties[identName] = identValue
	}

	p.expect(lexer.CLOSE_BRACKET)
	return ast.HeadersStmt{
		Properties: properties,
	}
}

func parseBodyStmt(p *parser) ast.Stmt {
	p.expect(lexer.BODY)
	if p.currentTokenKind() == lexer.ASSIGN {
		p.advance()

	}

	var properties = map[string]string{}
	p.expect(lexer.OPEN_BRACKET)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_BRACKET {
		var identName string

		identName = p.expect(lexer.IDENTIFIER).Value
		p.expectError(lexer.ASSIGN, "expected to find colon following property name")

		switch p.currentTokenKind() {
		case lexer.IDENTIFIER, lexer.STRING, lexer.NUMBER:
			properties[identName] = p.advance().Value
		default:
			panic("expected header valur after ':'")
		}

		// properties[identName] = identValue
	}

	p.expect(lexer.CLOSE_BRACKET)
	return ast.BodyStmt{
		Properties: properties,
	}
}

func parseBlockStmt(p *parser) ast.Stmt {
	p.expect(lexer.OPEN_BRACKET)
	body := []ast.Stmt{}
	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_BRACKET {
		body = append(body, parseStmt(p))
	}

	p.expect(lexer.CLOSE_BRACKET)
	return ast.BlockStmt{
		Body: body,
	}
}
