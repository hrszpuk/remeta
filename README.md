<p align="center">
    <img src="https://github.com/hrszpuk/remeta/assets/107559570/4fbf5d49-9701-43d9-b7d0-eab1f8e2ff8b" alt="remeta logo" width=200px>
</p>

<p align="center">
    Automatically generate Go package bindings for ReRect!
</p>

<p align="center">
<a href="./LICENSE.md"><img src="https://img.shields.io/badge/license-GPL-green.svg"></a>
<a href="https://github.com/hrszpuk"><img src="https://img.shields.io/github/followers/hrszpuk?style=social"></a>
<a href="https://github.com/hrszpuk/remeta/issues"><img src="https://img.shields.io/github/issues/hrszpuk/inimod"></a>
</p>

<p align="center">
    Remeta uses Go's parser and some metaprogramming magic to quickly generate easy to read bindings<br>
</p>

<p align="center">
    <a href="https://github.com/hrszpuk/remeta/blob/master/INSTALLATION.md">Installation</a>&nbsp;&nbsp;&nbsp;
    <a href="https://github.com/hrszpuk/remeta/blob/master/USER_GUIDE.md">User Guide</a>&nbsp;&nbsp;&nbsp;
    <a href="https://github.com/hrszpuk/remeta/blob/master/CONTRIBUTING.md">Contributing</a>&nbsp;&nbsp;&nbsp;
</p>


<table>
<tr>
<th>Go module</th>
<th>ReRect bindings</th>
</tr>
<tr>
<td>

```go
package test

type Vector3 struct {
	x int
	y int
	z int
}

func Vector3Add(v1, v2 Vector3) Vector3 {
	return Vector3{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
		z: v1.z + v2.z,
	}
}

func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3Add(v1, v2)
}

```

</td>
<td>

```go

package gopackages

import (
"bytespace.network/rerect/compunit"
"bytespace.network/rerect/eval_objects"
"bytespace.network/rerect/symbols"
_origin "bytespace.network/rerect/go_packages/test"
)

func LoadTest() {
	test := registerPackage("test")

	Vector3TypeSymbol := symbols.NewTypeSymbol(
		"Vector3",
		[]*symbols.TypeSymbol{},
		symbols.CON,
		0,
		nil,
	)

	Vector3Container := symbols.NewContainerSymbol(
		test, "Vector3", Vector3TypeSymbol,
	)

	Vector3Container.Fields = append(
		Vector3Container.Fields,
		symbols.NewFieldSymbol(
			Vector3Container, "x", 
			compunit.GlobalDataTypeRegister["int"],
		),
	)
	Vector3Container.Fields = append(
		Vector3Container.Fields,
		symbols.NewFieldSymbol(
			Vector3Container, "y", 
			compunit.GlobalDataTypeRegister["int"],
		),
	)
	Vector3Container.Fields = append(
		Vector3Container.Fields,
		symbols.NewFieldSymbol(
			Vector3Container, "z", 
			compunit.GlobalDataTypeRegister["int"],
		),
	)

	symbols.NewVMFunctionSymbol(
		test,
		"Vector3Add",
		compunit.GlobalDataTypeRegister["Vector3"],
		[]*symbols.ParameterSymbol{
			symbols.NewParameterSymbol(
				"v1",
				0,
				compunit.GlobalDataTypeRegister["Vector3"],
			),
		},
		Vector3Add,
	)
	symbols.NewVMFunctionSymbol(
		test,
		"Add",
		compunit.GlobalDataTypeRegister["Vector3"],
		[]*symbols.ParameterSymbol{
			symbols.NewParameterSymbol(
				"v2",
				0,
				compunit.GlobalDataTypeRegister["Vector3"],
			),
		},
		Add,
		)
}

func Vector3Add(args []any) any {
	v1 := args[0].(Vector3)
	v2 := args[0].(Vector3)
	return _origin.Vector3Add(v1, v2)
}

func Vector3_Add(instance any, args []any) any {
	v2 := args[0].(Vector3)
	return Add(v2)
}
```

</td>
</tr>
</table>
