package main

import "fmt"

type Generator struct {
}

type Param struct {
	Name string
	Type string
}

func GenerateRegisterFunction(name string, returnType string, params []string) string {
	fmt.Sprintf("registerFunction(")
}

func GenerateRegisterPackage(packageName string) string {
	return fmt.Sprintf("%s := registerPackage(\"%s\")\n", packageName, packageName)
}
