package main

import (
	"fmt"
	"math"
)

func main() {
	// 無名の関数リテラルをhypotに代入
	// hypotの型はfunc(float64, float64) float64
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	// コールバック
	// イベント処理、非同期処理

	// 戦略パターン
	// アルゴリズム

	// 関数クロージャ
	// 関数が自分の変数を覚えている

	adder := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}
	adder1 := adder()
	adder2 := adder()
	fmt.Println(adder1(1))
	fmt.Println(adder2(1))
}