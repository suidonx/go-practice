package main

import "fmt"

type MyInt int

func (m MyInt) String() string {
	return "hoge"
}

func main() {
	var m MyInt = 3
	fmt.Println(m) // hogeと出力させるように修正せよ。ただしmain関数に変更を加えないこと。
}

/*
Stringerインタフェース
String()を持つ
fmt.Printlnを呼ぶと自動的に呼ばれる

ベストプラクティスは
人間が読みやすい表示
	オーバーライドする
無限ループを避ける
	String()の中で、%vを使うと自動的にStringerインタフェースを呼び出すため。
nilレシーバ対応
	エラーハンドリングする
*/