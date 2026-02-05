package main

import (
	"client_http/lexer"
	"os"

	"github.com/sanity-io/litter"
)

func main() {
	data, err := os.ReadFile("./teste.aeon")
	if err != nil {
		panic(err.Error())
	}

	tokens := lexer.Tokenize(string(data))
	litter.Dump(tokens)
}
