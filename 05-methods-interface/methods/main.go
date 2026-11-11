package main

import (
	"fmt"
	"math"
)

// 構造体の定義
// Goにはクラスがないので、構造体がそれに相当する
type User struct {
	FirstName string
	LastName  string
}

// メソッドの定義
// レシーバ引数を持つ関数
func (u User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Step 1: 値レシーバで変更しても、呼び出し元は変わらない
func (v Vertex) ScaleValue(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Step 1 / 3: ポインタレシーバなら元の値が変わる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Step 2: 大きい構造体。値レシーバだと呼び出しのたびに丸ごとコピーされる
type Profile struct {
	Name string
	Bio  [1024]byte
}

func (p Profile) NameValue() string {
	return p.Name
}

func (p *Profile) NamePtr() string {
	return p.Name
}

func main() {
	user := User{
		FirstName: "Taro",
		LastName:  "Yamada",
	}
	// メソッドの呼び出し
	// ドット記法で呼び出す
	fmt.Println(user.FullName())

	v := Vertex{3, 4}
	fmt.Println(v.Abs())     // 5
	fmt.Println(AbsFunc(v)) // 5

	// C-typeのような対称性
	p := &Vertex{3, 4}
	fmt.Println(p.Abs())      // ポインタで値レシーバを呼び出す(*pを渡す)
	fmt.Println(AbsFunc(*p)) // 5

	// --- フローチャート Step 1: 変更するならポインタ ---
	a := Vertex{3, 4}
	a.ScaleValue(10)
	fmt.Println("値レシーバの Scale 後:", a) // {3 4}

	b := Vertex{3, 4}
	b.Scale(10)
	fmt.Println("ポインタレシーバの Scale 後:", b) // {30 40}

	// --- Step 2: 大きいならポインタ（コピー回避。結果は同じ）---
	profile := Profile{Name: "Taro"}
	fmt.Println("値レシーバ:", profile.NameValue())
	fmt.Println("ポインタレシーバ:", profile.NamePtr())

	// --- Step 3: 他がポインタなら揃える（Scale があるので Abs もポインタ向き）---
}
