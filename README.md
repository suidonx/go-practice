# go-practice

Go の練習用リポジトリ。

## 章（学習順）

1. `01-basics-functions` — パッケージ、import、関数など
2. `02-variables` — 変数、型、ゼロ値、型変換
3. `03-flow-control` — for、if など
4. `04-pointers-slices` — ポインタ、スライス
5. `05-methods-interface` — メソッド、インターフェース
6. `06-goroutine` — goroutine
7. `07-package` — 標準ライブラリ（scanner、strings など）
8. `08-testing` — テスト
9. `09-http-api` — HTTP API
10. `10-database-sql` — database/sql
11. `11-jira-clone` — Gin / GORM の層分け API

## 初級課題

`beginner-exercises/` — 章とは別にまとめた演習。

| 番号 | 内容 |
|---|---|
| `001`〜`006` | 小さな CLI 演習 |
| `007` | ログ JSON を PostgreSQL へ取り込む |

`007` は Docker で DB を立ててから実行する。手順は [`beginner-exercises/007/README.md`](beginner-exercises/007/README.md)。

## 使い方

各トピックフォルダで `main.go` を実行する。

```bash
cd 02-variables/zeros
go run main.go
```
