package ast

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

type MethodStmt struct {
	Value Expr // or string if you want to simplify
}

func (m MethodStmt) stmt() {}

// headers section
type HeadersStmt struct {
	Properties map[string]string
}

func (m HeadersStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) stmt() {}
