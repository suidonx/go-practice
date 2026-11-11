package main

import "fmt"

func main() {
	numbers := []interface{}{1, "2", 10, "11"}

	// 型 Switch
	for _, number := range numbers {
		switch v := number.(type) {
		case int:
			fmt.Printf("%02d\n", v)
		case string:
			fmt.Printf("%02s\n", v)
		}
	}
}