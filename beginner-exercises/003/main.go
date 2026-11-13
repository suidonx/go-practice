// type MyIntSlice []intでIntスライスと互換のある独自型を作り、重複を排除するメソッドを実装する。

package main

import "fmt"
import "encoding/json"

type MyIntSlice []int

func (p *MyIntSlice) Unique() []int {
	m := make(map[int]bool) // m := map[int]bool{}と同じ
	s := []int{}
	for _, v := range *p {
		if !m[v] {
			m[v] = true
			s = append(s, v)
		}
	}
	return s
}

func main () {
	m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Printf("%v\n", m.Unique()) // [1, 2, 3, 4, 5]
	jsonData, err := json.Marshal(m.Unique())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}