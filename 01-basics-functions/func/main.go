package main

import "fmt"

// 乗算
func multiple(x int, y int) int {
	return x * y
}

func addVerbose(x, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}

func rectangleArea(width, height int) int {
	return width * height
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(multiple(42, 13))
}