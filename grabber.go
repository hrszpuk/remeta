package main

import (
	"fmt"
	"go/ast"
)

type FuncData struct {
	Name string
	Type ast.FuncType
}

type Grabber struct {
	Functions []FuncData
}

func (g Grabber) GrabAll(file *ast.File) {
	ast.Walk(g, file)
}

func (g Grabber) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.FuncDecl:
		fmt.Printf("%v\n", d)
		f := FuncData{d.Name.Name, *d.Type}
		g.Functions = append(g.Functions, f)
		fmt.Println(g.Functions)
	}

	return g
}
