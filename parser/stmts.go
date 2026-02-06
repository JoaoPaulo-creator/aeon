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
