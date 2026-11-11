package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)
func main() {
	// DSN: Data Source Name
	dsn := "postgres://user:password@localhost:5432/go_practice?sslmode=disable"
	// データベースに接続プールを用意する。接続プールは、複数の接続を管理するための機能。
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// 接続プールを閉じる
	defer db.Close()

	// 接続プールをpingする。pingは、接続プールが正常に接続されているかを確認するためのメソッド。
	if err := db.Ping(); err != nil {
		log.Fatalf("ping failed: %v", err)
	}
	log.Println("connected")

}
