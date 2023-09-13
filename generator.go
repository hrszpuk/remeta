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

func (g *Generator) Generate() {
	// 0. Package name
	g.OutputSource = g.GeneratePackageDeclaration()
	g.OutputSource += "\n\n"

	// 1. Imports
	//g.OutputSource += g.GenerateImports()
	//g.OutputSource += "\n\n"

	// 2. Load function
	g.OutputSource += g.GenerateLoadFunc()

	// 3. Package Registration
	g.OutputSource += "\t" + g.GenerateRegisterPackage() + "\n"

	// 4. Function Registration
	for _, fn := range g.Functions {
		g.OutputSource += "\t" + g.GenerateFunctionSymbol(&fn) + "\n"
	}

	g.OutputSource += "}\n\n"

	// 5. Function Implementation
	for _, fn := range g.Functions {
		g.OutputSource += g.GenerateFunctionImplementation(&fn) + "\n\n"
	}
}

func (g *Generator) GeneratePackageDeclaration() string {
	return fmt.Sprintf("package %s", g.PackageName)
}

func (g *Generator) GenerateImports() string {
	output := "import (\n"

	for _, imp := range g.Imports {
		output += "\t\"" + imp.Path.Value + "\"\n"
	}

	return output + ")"
}

// GenerateLoadFunc This doesn't generate an ending bracket. Make sure to close!
func (g *Generator) GenerateLoadFunc() string {
	return fmt.Sprintf("func Load%s() {\n", g.PackageName)
}

func (g *Generator) GenerateRegisterFunction(packageName string, function *ast.FuncDecl) string {
	return fmt.Sprintf("registerFunction(\"%s\", %s)", g.PackageName, g.GenerateFunctionSymbol(function))
}

func (g *Generator) GenerateFunctionSymbol(function *ast.FuncDecl) string {
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

func (g *Generator) GenerateParameterSymbol(field *ast.Field, index int) string {

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

func (g *Generator) GenerateParameterSymbolSlice(function *ast.FuncDecl) string {
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

func (g *Generator) GenerateRegisterPackage() string {
	return fmt.Sprintf("%s := registerPackage(\"%s\")\n", g.PackageName, g.PackageName)
}

func (g *Generator) GenerateGlobalDataTypeRegister(typeName string) string {
	return fmt.Sprintf("compunit.GlobalDataTypeRegister[\"%s\"]", typeName)
}
