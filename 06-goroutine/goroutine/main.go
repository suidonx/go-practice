package main

import (
	"fmt"
	_ "time"
)

func sum(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total // チャネルに結果を送信 // 結果がcに届くまでここで待機
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0} // スライス

	c := make(chan int) // チャネルを作成
	go sum(s[:3], c) // ゴルーチンを起動して、前半3つをsum
	go sum(s[3:], c) // ゴルーチンを起動して、後半3つをsum
	x, y := <-c, <-c // チャネルから結果を受け取る(届くまでここで待機)
	fmt.Println(x, y, x+y)

	// x := 1
	// go fmt.Println(x)
	// x = 2

	// y, z := 10, 20
	// go sum(y, z)
	// y = 100

	// time.Sleep(1 * time.Second)
}

// func sum(y, z int) {
// 	fmt.Println(y + z)
// }