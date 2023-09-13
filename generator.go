package main

import (
	"fmt"
	"go/ast"
)

// Generator Takes stored nodes and uses them to generate Go code
type Generator struct {
	// Meta data
	PackageName    string
	OutputFileName string
	OutputSource   string

	// Parser data
	Functions []ast.FuncDecl
	Variables []ast.GenDecl
	Imports   []ast.ImportSpec
}

func NewGenerator(g *Grabber, name string) *Generator {
	gen := new(Generator)
	gen.Functions = g.Functions
	gen.Variables = g.Variables
	gen.PackageName = name
	gen.OutputFileName = name + ".go"
	return gen
}

}

func GenerateFunctionSymbol(packageName string, function *ast.FuncDecl) string {
	// ReRect only currently allows for a single return type so anything returning multiple values will probably be ignored for now.

	returnType := ""
	if function.Type.Results.NumFields() == 0 {
		returnType = "void"
	} else {
		ident, ok := function.Type.Results.List[0].Type.(*ast.Ident)
		if ok {
			returnType = ident.Name
		}
	}

	return fmt.Sprintf(
		"symbols.NewVMFunctionSymbol(%s, \"%s\", %s, %s, %s)",
		packageName,
		function.Name.String(),
		GenerateGlobalDataTypeRegister(returnType),
		GenerateParameterSymbolSlice(function),
		function.Name.String(),
	)
}

func GenerateParameterSymbol(field *ast.Field, index int) string {

	returnType := ""
	if t, ok := field.Type.(*ast.Ident); ok {
		returnType = t.Name
	}

	return fmt.Sprintf(
		"symbols.NewParameterSymbol(\"%s\", %d, %s)",
		field.Names[0],
		index,
		GenerateGlobalDataTypeRegister(returnType),
	)
}

func GenerateParameterSymbolSlice(function *ast.FuncDecl) string {
	output := "[]*symbols.ParameterSymbol{"
	for i, s := range function.Type.Params.List {
		output += GenerateParameterSymbol(s, i) + ", "
	}
	if function.Type.Params.NumFields() > 0 {
		output = output[:len(output)-2]
	}
	output += "}"
	return output
}

func GenerateRegisterPackage(packageName string) string {
	return fmt.Sprintf("%s := registerPackage(\"%s\")\n", packageName, packageName)
}

func GenerateGlobalDataTypeRegister(typeName string) string {
	return fmt.Sprintf("compunit.GlobalDataTypeRegister[\"%s\"]", typeName)
}
