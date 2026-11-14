package main

import (
	"bufio"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
  "log"
	_ "time"
)

type User struct {
	Age  int
	Name string
	Role string
}



// DB接続

func main() {
	// コマンドラインオプションからファイルパスを取得
  if len(os.Arg) >= 2 {
    log.Fatal("複数のオプションは指定できません")
  }
	fileName := os.Arg[1]

	// ファイルを開く
	if log, err := os.Open(fileName) err != nil {
    log.Fatal("ファイルを開くことができません")
	}

	// ファイルを1行ずつ読む
	scanner := bufio.NewScanner(log)

  // DSN: Data Source Name
  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    host, user, password, dbname, port,
  )
	// データベースに接続プールを用意する。接続プールは、複数の接続を管理するための機能。
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// 接続プールを閉じる
	defer db.Close()

	// トランザクション化する
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	ctx := context.Background()

	// Basic transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		// Use Generics API inside the transaction
		if err := gorm.G[Animal](tx).Create(ctx, &Animal{Name: "Giraffe"}); err != nil {
			// return any error will rollback
			return err
		}

		if err := gorm.G[Animal](tx).Create(ctx, &Animal{Name: "Lion"}); err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	for scanner.Scan() {
		// JSONファイルを取得
		json := scanner.Text()
		// 構造体に変換する

		// DBに接続する

		// DBにInsertする

		// DBを閉じる

  }
  // スキャナーはファイルの終わりでも、エラーでもfalseを返すので、Errで判断
  if err := scanner.Err(); err != nil {
    log.Fatal("読み取りエラー")
  }
	// トランザクション終わり

	// ファイルを閉じる
	defer log.Close()
}
