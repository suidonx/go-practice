package main

import "fmt"

func main() {
	// for basic
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum =", sum)

	// count down
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("count down done")

	// 配列・スライスのlinear search
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 省略記法
	s := 1
	for ; s < 1000; s *= 2 {
		fmt.Println(s)
	}

	// 後処理を省略
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	// while風
	sum = 1
	for ; sum < 1000; {
		sum += sum
		fmt.Println("sum =", sum)
	}

	// count up
	// セミコロンを含めない
	count := 0
	for count < 10 {
		fmt.Println("count =", count)
		count++
	}

	// infinite loop
	countInfinite := 0
	for {
		countInfinite++
		if countInfinite > 10 {
			break
		}
		fmt.Println("countInfinite =", countInfinite)
	}

	// continue
	for i := 0; i < 10; i++ {
		// 偶数をスキップ
		if i % 2 == 0 {
			continue
		}
		fmt.Println("i =", i)
	}
}
