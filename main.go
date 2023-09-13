package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "test/test.go", nil, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	g := new(Grabber)
	g.GrabAll(f)

	generator := NewGenerator(g, "test")
	generator.Generate()
	fmt.Println(generator.OutputSource)
}
