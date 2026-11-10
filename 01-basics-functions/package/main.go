package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Printf("My number is %d.\n", rand.Intn(10))

	name := "nico"
	fmt.Printf("name is %T\n", name)
}
