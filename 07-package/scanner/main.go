package main

import (
	"fmt"
	"bufio" // バッファ付きの読み取り
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file) // ファイルを読み取るスキャナを作成
	for scanner.Scan() { // ファイルを一行ずつ読み取る
		fmt.Println(scanner.Text()) // 読み取った行を出力
	}
	// スキャナはデフォルトで行ごとに読み取る、単語ごと、rune文字ごと、バイトごとにすることが可能
}