package main

import (
	"fmt"
	"math"
)

func main() {
	// if basic
	age := 20
	if age >= 20 {
		fmt.Println("adult")
	}

	// short statement
	// forループの初期化と同じ仕組み
	// if文の中で変数が有効
	// 変数の寿命を短くするのがGoの思想
	if age := 20; age >= 20 {
		fmt.Println("age =", age)
	}

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	// 実用的なパターン
	// if err := doSomething(); err != nil {
	// 	fmt.Println("error =", err)
	// }

	// 計算結果の判定
	// if result := calculate(); result > threshold {
	// 	// 閾値を越えた場合の処理
	// }

	// if else
	if age >= 20 {
		fmt.Println("adult")
	} else {
		fmt.Println("child")
	}

	// if else if else
	if age >= 20 {
		fmt.Println("adult")
	} else if age >= 10 {
		fmt.Println("child")
	} else {
		fmt.Println("baby")
	}

	// Goのearly return
	// 関数の途中でreturnを使って処理を抜ける

	// 自動セミコロンルール
	// if文の中では自動的にセミコロンが挿入される

	// switch basic
	// 短絡評価 short-circuit evaluation
	// caseが一致したらその後のcaseは評価されない
	// 重い処理は後ろに書く
	switch day := "Monday"; day {
	case "Monkey":
		fmt.Println("Monday")
	case "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Holiday")
	}

	// defer
	// deferは関数の実行を遅らせる
	// 関数の実行が終わった後に実行される
	// ファイルのクローズ、ロックの解放、DB接続のクローズ
	// 書き忘れ防止のために使う

	defer fmt.Println("world")
	fmt.Println("hello")

	i := 0
	defer fmt.Println("i =", i)
	i++
	fmt.Println("i =", i)

	// deferのスタック
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
