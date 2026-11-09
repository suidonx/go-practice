package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func divmod(x, y int) (int, int) {
	quotient := x / y // 商
	remainder := x % y // 余り
	return quotient, remainder
}

// 戻り値を無視する
func ignore() {
	quotient, _ := divmod(7, 3) // 余りを無視
	fmt.Println(quotient)
}

// 戻り値を無視しない
func ignore2() {
	quotient, remainder := divmod(7, 3)
	fmt.Println(quotient, remainder)
}

// 名前付き戻り値
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 名前付き戻り値通常版
func split2(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	a, b := swap("hello", "world") // 多重代入
	fmt.Println(a, b)
	ignore()

	split(17)
	split2(17)
}