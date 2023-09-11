package main

import (
	"fmt"
	"go/ast"
)

func GenerateRegisterFunction(packageName string, function *ast.FuncDecl) string {
	return fmt.Sprintf("registerFunction(\"%s\", %s)", packageName, GenerateFunctionSymbol(packageName, function))
}

func GenerateRegisterPackage(packageName string) string {
	return fmt.Sprintf("%s := registerPackage(\"%s\")\n", packageName, packageName)
}
