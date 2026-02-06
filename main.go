package main

import (
	"client_http/lexer"
	"client_http/parser"
	"os"

	"github.com/sanity-io/litter"
)

func main() {
	data, err := os.ReadFile("./teste3.aeon")
	if err != nil {
		panic(err.Error())
	}

	tokens := lexer.Tokenize(string(data))
	ast := parser.Parse(tokens)

	litter.Dump(ast)
}
