package main

import "fmt"

func main() {
	var a = 10
	var b = -5
	var c = 1.41421356

	var d = "Hello, World"
	var e = true
	var f = 'A'

	// 型推論
	fmt.Printf("a: %T, b: %T, c: %T\n", a, b, c)
	fmt.Printf("d: %T, e: %T, f: %T\n", d, e, f)
}