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
	Structs   []ast.TypeSpec
	Imports   []ast.ImportSpec
}

func NewGenerator(g *Grabber, name string) *Generator {
	gen := new(Generator)
	gen.Functions = g.Functions
	gen.Imports = g.Imports
	gen.Structs = g.Structs
	gen.PackageName = name
	gen.OutputFileName = name + ".go"
	return gen
}

func (g *Generator) Generate() {
	// 0. Package name
	g.OutputSource = g.GeneratePackageDeclaration()
	g.OutputSource += "\n\n"

	// 1. Imports
	g.OutputSource += g.GenerateImports()
	g.OutputSource += "\n\n"

	// 2. Load function
	g.OutputSource += g.GenerateLoadFunc()

	// 3. Package Registration
	g.OutputSource += "\t" + g.GenerateRegisterPackage() + "\n\n"

	for _, con := range g.Structs {
		name := con.Name.String()
		containerName := name + "Container"
		typeSymbolName := name + "TypeSymbol"
		g.OutputSource += fmt.Sprintf("\t%s := %s\n", typeSymbolName, g.GenerateTypeSymbol(name, "symbols.CONT"))
		g.OutputSource += fmt.Sprintf("\t%s := %s\n\n", containerName, g.GenerateContainerSymbol(g.PackageName, &con, typeSymbolName))

		for _, field := range con.Type.(*ast.StructType).Fields.List {
			g.OutputSource += fmt.Sprintf(
				"\t%s.Fields = append(%s.Fields, %s)\n",
				containerName,
				containerName,
				g.GenerateFieldSymbol(containerName, field.Names[0].String(), field.Type.(*ast.Ident).Name),
			)
		}

		g.OutputSource += "\n"
	}

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
	return "package gopackages"
}

func (g *Generator) GenerateImports() string {
	output := "import (\n"

	for _, imp := range g.Imports {
		output += "\t" + imp.Path.Value + "\n"
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

func (g *Generator) GenerateRegisterContainer(packageName string, containerName string) string {
	return fmt.Sprintf("registerContainer(\"%s\", %s)", g.PackageName, containerName)
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
		g.PackageName,
		function.Name.String(),
		g.GenerateGlobalDataTypeRegister(returnType),
		g.GenerateParameterSymbolSlice(function),
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
		g.GenerateGlobalDataTypeRegister(returnType),
	)
}

func (g *Generator) GenerateParameterSymbolSlice(function *ast.FuncDecl) string {
	output := "[]*symbols.ParameterSymbol{"
	for i, s := range function.Type.Params.List {
		output += g.GenerateParameterSymbol(s, i) + ", "
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

func (g *Generator) GenerateFunctionImplementation(function *ast.FuncDecl) string {
	out := fmt.Sprintf("func %s(args []any) any {\n", function.Name.String())

	arguments := make([]string, 0)

	if function.Type.Params.NumFields() > 0 {
		for i, p := range function.Type.Params.List {
			out += fmt.Sprintf(
				"\t%s := args[%d].(%s)\n",
				p.Names[0].String(),
				i,
				p.Type,
			)
			arguments = append(arguments, p.Names[0].String())
		}
	}

	GenerateFunctionCall := func(name string, args []string) (res string) {
		res += function.Name.String() + "("
		for i, arg := range arguments {
			res += arg
			if i != len(arguments)-1 {
				res += ", "
			}
		}
		return res + ")\n"
	}

	if function.Type.Results == nil {
		out += "\t" + GenerateFunctionCall(function.Name.String(), arguments) + "\treturn nil\n"
	} else {
		out += "\treturn " + GenerateFunctionCall(function.Name.String(), arguments)
	}

	out += "}"

	return out
}

// typ = "symbols.CONT" // for Container
func (g *Generator) GenerateTypeSymbol(name string, typ string) string {
	return fmt.Sprintf("symbols.NewTypeSymbol(\"%s\", []*symbols.TypeSymbol{}, %s, 0, nil)", name, typ)
}

func (g *Generator) GenerateContainerSymbol(packageName string, st *ast.TypeSpec, typeSymbolName string) string {
	out := "container " + st.Name.String() + "{\n"

	for _, field := range st.Type.(*ast.StructType).Fields.List {
		out += "\t" + field.Names[0].String() + " " + field.Type.(*ast.Ident).Name + "\n"
	}

	out += "}"

	return fmt.Sprintf("symbols.NewContainerSymbol(%s, \"%s\", %s)", packageName, st.Name.String(), typeSymbolName)
}

func (g *Generator) GenerateFieldSymbol(containerName, fieldName, dataType string) string {
	return fmt.Sprintf("symbols.NewFieldSymbol(%s, \"%s\", %s)", containerName, fieldName, g.GenerateGlobalDataTypeRegister(dataType))
}
