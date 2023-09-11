package main

import (
	"fmt"
	"go/ast"
)

func GenerateRegisterFunction(packageName string, function *ast.FuncDecl) string {
	return fmt.Sprintf("registerFunction(\"%s\", %s)", packageName, GenerateFunctionSymbol(packageName, function))
}

func GenerateFunctionSymbol(packageName string, function *ast.FuncDecl) string {
	// ReRect only currently allows for a single return type so anything returning multiple values will probably be ignored for now.

	returnType := ""
	if function.Type.Results.NumFields() == 0 {
		returnType = "void"
	} else {
		returnType = "" //function.Type.Results.List[0].Type
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
	return fmt.Sprintf(
		"symbols.NewParameterSymbol(\"%s\", %d, %s)",
		field.Names[0],
		index,
		GenerateGlobalDataTypeRegister(""),
	)
}


func GenerateRegisterPackage(packageName string) string {
	return fmt.Sprintf("%s := registerPackage(\"%s\")\n", packageName, packageName)
}

func GenerateGlobalDataTypeRegister(typeName string) string {
	return fmt.Sprintf("compunit.GlobalDataTypeRegister[\"%s\"]", typeName)
}
