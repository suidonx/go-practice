package main

import "fmt"

func main() {
	numbers := []any{1, "2", 10, "11"}

	// 型 Switch
	for _, number := range numbers {
		switch number := number.(type) {
		case int:
			fmt.Printf("%02d\n", number)
		case string:
			fmt.Printf("%02s\n", number)
		}
	}
}