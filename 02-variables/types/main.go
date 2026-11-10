package main

import "fmt"

func main() {
	// boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of isLoggedIn is %T\n", isLoggedIn)

	// integer
	var age int = 20
	fmt.Println(age)
	fmt.Printf("Type of age is %T\n", age)

	// float
	var price float64 = 10.5
	fmt.Println(price)
	fmt.Printf("Type of price is %T\n", price)
}