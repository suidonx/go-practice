package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8}
	for i, v := range pow {
		fmt.Printf("pow[%d] = %d\n", i, v)
	}

	sum := 0
	for _, value := range pow {
		sum += value
		fmt.Println("sum:", sum)
	}

	// インデックスのみ
	for i := range pow {
		fmt.Println("index:", i)
	}

	// 値のみ
	for _, value := range pow {
		fmt.Println("value:", value)
	}

	// インデックスと値
	for i, value := range pow {
		fmt.Printf("pow[%d] = %d\n", i, value)
	}

	// 実行回数だけ
	// 要素数の数だけ実行
	for range 3 {
		fmt.Println("range 3")
	}
}