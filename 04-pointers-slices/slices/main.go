package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // 配列
	s := primes[1:4]
	s[0] = 100
	fmt.Println("スライス:", s) // [100, 5, 7]
	fmt.Println("配列:", primes) // [2, 100, 5, 7, 11, 13]

	a := []int{1, 2, 3, 4, 5} // スライス、裏に配列がある。
	double(a)
	fmt.Println("after double:", a)

	b := [5]int{1, 2, 3, 4, 5}
	doubleArray(b)
	fmt.Println("after double:", b)

	// cap
	ss := []int{1, 2, 3, 4, 5}
	// 0番目から全て見える状態、容量は5
	fmt.Println("cap:", cap(ss)) // 5

	ss = ss[1:3]
	// 左から1個捨てたので、容量は4
	// 3番目の手前まで見えるので、長さは2
	fmt.Println("cap:", cap(ss)) // 4

	ss = ss[:4]
	// 今見えているssは[2, 3]
	// そこから見える数を4個にする。[2, 3, 4, 5]
	fmt.Println("cap:", cap(ss)) // 4

	// nilスライス
	var slice []int
	fmt.Println("nilスライス:", slice) // []
	fmt.Println("len:", len(slice)) // 0
	fmt.Println("cap:", cap(slice)) // 0

	slice = append(slice, 1)
	fmt.Println("nilスライス:", slice) // [1]
	fmt.Println("len:", len(slice)) // 1
	fmt.Println("cap:", cap(slice)) // 1

	// make
	slice = make([]int, 5, 10) // 長さ5, 容量10のスライスを作る
	// 裏配列をcap個確保する。
	// 裏配列は全て0で初期化
	fmt.Printf("make([]int, 5): len=%d, cap=%d\n", len(slice), cap(slice))

	slice = make([]int, 0, 10) // 長さ0, 容量10のスライスを作る
	// 中身は空だけど、100個分の領域を確保する。
	fmt.Printf("make([]int, 0, 10): len=%d, cap=%d\n", len(slice), cap(slice))

	// 二次元スライス
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("行列:", matrix)

}

func double(a []int) {
	for i := range a {
		a[i] = a[i] * 2
	}
}

// 配列バージョン
func doubleArray(a [5]int) {
	for i := range a {
		a[i] = a[i] * 2
	}
}