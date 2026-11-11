package main

import "fmt"

func describe(i any) {
	switch v := i.(type) {
	case string:
		fmt.Printf("string: %s\n", v)
	case int:
		fmt.Printf("int: %d\n", v)
	default:
		fmt.Printf("unknown: %v\n", v)
	}
}

func main() {
	var i any = "hello"
	describe(i)
	describe(10)
	describe("hello")
	describe(10.0)
}
