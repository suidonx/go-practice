package main

import "fmt"
/*
mapのvalueからkeyを取得する問題。
引数にmapとvalueを受け取り、対応するkeyがあればkeyを返却する関数を実装する。
mapの型はkeyをint, valueをstringにすること。もし対応するkeyがなければerrorを返すこと。
*/

func main() {
  m := map[int]string{
    1: "01",
    2: "02",
    3: "03",
  }
  key, err := findKeyByValue(m, "03") // key→3, err→nil
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Println(key)
  }
  
  key, err = findKeyByValue(m, "05") // key→0にすること(初期値なので), errはある
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Println(key)
  }

}

// map逆引き
func findKeyByValue(m map[int]string, value string) (int, error) {
  for key, val := range m {
    if value == val {
	  return key, nil
	}
  }
  return 0, fmt.Errorf("%s に対応する key はありません", value)
}
