# 007 ログ取り込み

このプログラムは、`sample.log` に書かれたログを読み、その中のユーザ情報を PostgreSQL に保存します。

`sample.log` は、1行につき JSON が1件です。各行の `user`（年齢・名前・役割）が `users` テーブルの1行になります。

## 必要なもの

- Docker（データベースをコンテナで起動します）
- Go
- ターミナルのカレントディレクトリを、このフォルダ（`beginner-exercises/007`）にする

## 手順

### 1. データベースを起動する

```bash
docker compose up -d
```

`-d` は、コンテナを裏側で動かし続ける指定です。初回はイメージのダウンロードで少し時間がかかります。

起動できたか確認します。

```bash
docker compose ps
```

`postgres` の状態が running なら次へ進みます。接続先は自分のマシンの `localhost:5432` です。

### 2. 接続情報をターミナルに渡す

Go は `.env` ファイルを自分では読みません。起動したプログラムが見るのは、**そのときターミナルに設定されている環境変数**だけです。

`.env` には DB のホスト・ポート・ユーザ・パスワード・DB 名が書いてあります。次の3行で、その内容を今のターミナルに読み込みます。

```bash
set -a
source .env
set +a
```

Docker で動かしている Postgres と、`go run` するプログラムは別物です。コンテナ側の設定はステップ 1 で入りますが、手元の Go には入らないので、このステップが必要です。

すでに同じ内容を `export` 済みのターミナルなら、飛ばして構いません。

### 3. プログラムを実行する

```bash
go run main.go sample.log
```

`sample.log` は「どのファイルを取り込むか」で、コマンドの引数として渡します。

成功すると次のように表示されます。

```text
インポートが完了しました
```

実行中にプログラムがやっていることは次のとおりです。

1. `users` テーブルが無ければ作る
2. `sample.log` を1行ずつ読む
3. 各行の `user` をデータベースに INSERT する

## 結果の見方

`sample.log` は5行あるので、初めて実行したときは `users` に5件入ります。もう一度同じコマンドを実行すると、同じ人がまた5件追加されます（重複チェックはありません）。

### コマンドで見る

コンテナの中の `psql`（Postgres の対話クライアント）でつなぎます。

```bash
docker compose exec postgres psql -U user -d 007_go_practice_db
```

つながるとプロンプトが `007_go_practice_db=#` になります。テーブルと中身を見ます。

```sql
\dt
SELECT * FROM users;
```

初回成功時は次のように5件出ます。

```text
 id | age |  name   |  role
----+-----+---------+---------
  1 |  22 | tarou   | student
  2 |  23 | zirou   | student
  3 |  24 | saburou | student
  4 |  25 | mike    | mentor
  5 |  26 | make    | mentor
```

抜けるときは `\q` です。1発で見るなら:

```bash
docker compose exec postgres psql -U user -d 007_go_practice_db -c 'SELECT * FROM users;'
```

`users` が無いときは、まだステップ 3 を実行していません。テーブルはプログラムの AutoMigrate で作られます。

### GUI（TablePlus など）で見る

コンテナの中に入るのではなく、ホストの `localhost:5432` に通常の Postgres としてつなぎます。値は `.env` と同じです。

| 項目 | 値 |
|---|---|
| Host | `localhost` |
| Port | `5432` |
| User | `user` |
| Password | `password` |
| Database | `007_go_practice_db` |
| SSL | オフ |

左に `users` が見えたら、中身はコマンドの `SELECT` と同じ5件です。

## 止めるとき

```bash
docker compose down
```

データベースの中身は Docker の volume（`postgres_data`）に残ります。データごと消すときは次です。

```bash
docker compose down -v
```

## 実装メモ

取り込み全体を一つのトランザクションにする話は [`transaction.md`](transaction.md)。
