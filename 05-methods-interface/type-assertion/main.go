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
	s := i.(string)
	fmt.Println(s)
	n, ok := i.(int)
	fmt.Println(n, ok) // 型があってないので、okがfalse、intのゼロ値が入る。この場合。

	describe(i)
	describe(10)
	describe("hello")
	describe(10.0)
}