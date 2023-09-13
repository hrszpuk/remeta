package main

import (
	"go/ast"
)

// Grabber "Grabs" nodes from the Go package and stores them
type Grabber struct {
	Functions []ast.FuncDecl
	Structs   []ast.TypeSpec
	Imports   []ast.ImportSpec
}

func (g *Grabber) GrabAll(file *ast.File) {
	g.Functions = make([]ast.FuncDecl, 0)
	g.Imports = make([]ast.ImportSpec, 0)
	g.Structs = make([]ast.TypeSpec, 0)
	ast.Walk(g, file)
}

func (g *Grabber) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.FuncDecl:
		g.Functions = append(g.Functions, *d)
	case *ast.ImportSpec:
		g.Imports = append(g.Imports, *d)
	case *ast.TypeSpec:
		_, ok := d.Type.(*ast.StructType)
		if ok {
			g.Structs = append(g.Structs, *d)
		}
	}

	return g
}
