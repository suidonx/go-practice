## テスト

標準ライブラリ
テストファイルの命名規則
    _test.go
    go testで実行

テストケース名
    valid email
    empty string
    nil pointer

テストケース構成
    input 入力
    want 期待値
    got 実際の値
    エラー有無

アサーション
    1つのテストケース＝1つのアサーション

テストカバレッジ
    正常系・異常系・エッジケース
    go test -cover

`go test`が落ちた

リポジトリにgo.modがないため。
Goモジュールの宣言が必要。
ルートを知るため。
どの依存ルールの上にいるのか知るため。