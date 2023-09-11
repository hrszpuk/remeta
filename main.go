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

	var g Grabber
	g.GrabAll(f)

	fmt.Println(g.Functions)
}
