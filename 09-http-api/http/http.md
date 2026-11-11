## HTTP API

標準ライブラリが充実
net/http: HTTPサーバ
encoding/json: JSONシリアライズ
database/sql: DB接続

外部ライブラリなしで実装可能

## パフォーマンス

コンパイル言語で実行速度が速い
ゴルーチンで軽量並行処理
メモリ効率が良い

## デプロイ・運用

コンパイルするとバイナリ1つになりコンテナ化しやすい
実行に Go 本体は不要。OS向けの1ファイルを置けば動くので、Docker はバイナリを COPY するだけでよい。
    Python/Node.jsはランタイムそのものが必要
    ライブラリ
    ソースコードなども必要    

`go build`で、OS用の実行ファイルが1個できる

マルチプレクサ
    届いたリクエストを見て、どのハンドラに渡すかを決める
    デフォルトのマルチプレクサに"/"はHelloHandlerに渡すと登録する
    nilはデフォルトを使うという意味
```
http.HandleFunc("/", HelloHandler)
// サーバーを起動
log.Fatal(http.ListenAndServe(":8080", nil))
```

curlコマンド
    ターミナルからHTTPリクエストを送るコマンド