package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	m["cherry"] = 3
	fmt.Println("map:", m)
}