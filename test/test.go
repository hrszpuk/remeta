package test

import (
	"fmt"
	"log"
)

var x = 1

func main() int {
	var y = x + 1
	fmt.Println("hello")
	log.Println("test")
	return y
}

func testFunc1() {}

func testFunc2(x int, s string) {}

func testFunc3(x int, s string) any {
	return nil
}
