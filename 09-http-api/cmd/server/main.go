package main

import (
	"fmt"
	"log"
	"net/http"
)

// ハンドラー関数
// リクエストを受け取り、レスポンスを返す関数
// wにレスポンスを書き込むことで、クライアントにレスポンスを返すことができる
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンスヘッダーを設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// ステータスコードを設定
	w.WriteHeader(http.StatusOK)
	// レスポンスボディを設定
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

// ハンドラーの構造体にDB/GORM, service, 設定、loggerなどを注入すると依存を減らす
// なぜ？ 依存を減らすことで、ハンドラーのテストがしやすくなる。
// 依存をフィールドに閉じ込める
type MyHandler struct {
	name string
}

// ServeHTTPメソッドを実装することで、http.Handlerインターフェースを満たす
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from " + h.name))
}

// リクエストハンドラー
// リクエストを受け取り、ログを出力する
// r.MethodはGET, POST, PUT, DELETEなどのHTTPメソッドを取得できる
// r.URL.Pathはリクエストのパスを取得できる
// r.URL.Queryはリクエストのクエリパラメータを取得できる
// r.Headerはリクエストのヘッダーを取得できる
// r.Bodyはリクエストのボディを取得できる
// r.Contextはリクエストのコンテキストを取得できる
// r.RemoteAddrはリクエストのリモートアドレスを取得できる
// r.UserAgentはリクエストのユーザーエージェントを取得できる
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: %s, Path: %s, Query: %v", r.Method, r.URL.Path, r.URL.Query())
	log.Printf("User-Agent: %s", r.Header.Get("User-Agent"))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"method": "%s", "path": "%s", "query": "%v"}`, r.Method, r.URL.Path, r.URL.Query())
}

func CreatedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created"))
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error": "invalid request"}`))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

func main() {
	// ハンドラー関数ではなく、型のインスタンスを作成して登録することで、依存を減らす。
	// http.Handle("/", &MyHandler{name: "MyServer"})
	// http.HandleFunc("/", RequestHandler)

	http.HandleFunc("/created", CreatedHandler)
	http.HandleFunc("/error", ErrorHandler)

	// サーバーを起動
	// nilはデフォルトのマルチプレクサーを使用
	// マルチプレクサーは、リクエストを受け取り、適切なハンドラー関数にルーティングする
	// ポート番号を指定
	// エラーが発生した場合はログを出力して終了
	log.Fatal(http.ListenAndServe(":8080", nil))
}