package main

import (
	"go/ast"
)

// Grabber "Grabs" nodes from the Go package and stores them
type Grabber struct {
	Functions []ast.FuncDecl
	Variables []ast.GenDecl
}

func (g *Grabber) GrabAll(file *ast.File) {
	g.Functions = make([]ast.FuncDecl, 0)
	ast.Walk(g, file)
}

func (g *Grabber) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.FuncDecl:
		g.Functions = append(g.Functions, *d)
	}

	return g
}
