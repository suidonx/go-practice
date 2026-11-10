package main

import "fmt"

func main() {
	var price float64 = 10.5
	fmt.Printf("%.2f\n", price)

	var x float64 = 10.5
	var y int = int(x)
	fmt.Printf("%.2f -> %d\n", x, y)

	// 型変換は明示的
	var i int = 42
	var f float64 = i

	// 異なる型同士の演算はエラー
	var result = x + y
}