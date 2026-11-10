package main

import "fmt"

func main() {
	// pointer basic
	// pointerはメモリの場所を示す
	// ポインタを住所とすると、家が値
	// ポインタを使うと、値を変更できる
	// ポインタを使うと、値を渡すのにコピーを作らない
	// ポインタを使うと、値を変更できる

	i := 42
	p := &i // &はポインタを取得、iのアドレスを取得
	fmt.Println("i =", i, "p =", p, "*p =", *p) // 値, ポインタ, ポインタの値
	// ポインタ渡し
	*p = 10
	fmt.Println("i =", i, "p =", p, "*p =", *p) // 値, ポインタ, ポインタの値

	// 関数とポインタ　値渡し vs ポインタ渡し
	// 値渡し：関数内では値が変更されない
	// ポインタ渡し：関数内では値が変更される
	i = 42
	doubleValue(i)
	fmt.Println("i =", i) // 42
	doubleValueByPointer(&i)
	fmt.Println("i =", i) // 84
}

func doubleValue(n int) {
	n = n * 2
}

func doubleValueByPointer(n *int) {
	*n = *n * 2
}