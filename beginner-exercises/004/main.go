package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []*User{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}

	for _, user := range users {
		(*user).Age = 44
	}

	for _, user := range users {
		fmt.Printf("%v\n", *user)
	}
	// fmt.Printf("%v", users) // どうなる？
}

/*
予想：
Userという構造体型で、Name, Ageフィールドがある
User型のスライスを初期化して変数usersに代入
{"tarou", 44},
{"zirou", 44},
{"itirou", 44},

答え：
rangeのuserはコピー
フィールド順だから、{name, age}の順でスライスに代入される
*/