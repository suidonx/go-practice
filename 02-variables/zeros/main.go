package main

import "fmt"

func main() {
	// ゼロ値
	var counter int
	counter++
	counter++
	counter++
	fmt.Println(counter)

	var message string
	message += "Hello, "
	message += "World"
	fmt.Println(message)
}