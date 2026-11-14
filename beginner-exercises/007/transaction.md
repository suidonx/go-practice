# 007 トランザクションメモ

`sample.log` の INSERT を **全部成功か全部失敗** にする仕組み。公式は [GORM Transactions（Traditional API）](https://gorm.io/docs/transactions.html)。

## 1. 何を保証するか

| 単位 | いつ確定するか |
|---|---|
| `Create` 1回（GORM のデフォルト） | その1行が成功したらすぐ COMMIT。3行目失敗でも 1・2 行目は残る |
| 今回の塊 | `Import` の for 全体。1件でも失敗（INSERT / Unmarshal / Scanner）したら ROLLBACK。全部 `nil` ならまとめて COMMIT |

塊に入れないもの: `os.Open`（DB ではない）。

## 2. 役割

```text
service  … 塊の範囲を決める。失敗したら return err
repo     … GORM を閉じ込める。Create と Transaction のラップ
GORM     … BEGIN / COMMIT / ROLLBACK。コールバックの戻り値だけ見る
```

service は `*gorm.DB` も `tx` も見ない。

repo は見た目同じ型が2つある。

| 変数 | `db` 欄 | 意味 |
|---|---|---|
| `s.repo` / `r` | 本番の接続 | main が渡した `*gorm.DB` |
| `fn` の `repo` | `tx` | BEGIN 済みの下書き |

`Create` のコードは一本。つながっている `r.db` が違うだけ。

## 3. なぜラップするか

`db.Create` を `Create` にしたのと同じで、**GORM を repository パッケージに閉じる**ため。

ラップしないと service は `s.repo.db.Transaction` と `tx.Create` が必要になる。`db` は小文字で他パッケージから見えない。見えたとしても層が崩れる。

## 4. なぜ `tx` と下書き repo を用意するか

GORM が渡すのは `tx *gorm.DB`。ドキュメントならここで `tx.Create` する。

007 の `fn` が欲しいのは `*UserRepository`（`repo.Create`）。`tx` をそのまま渡すと型が合わない。

```text
tx                         *gorm.DB（下書きの接続）
&UserRepository{db: tx}    *UserRepository（fn の引数 repo）
```

`&` は UserRepository 全体のアドレス。`&tx` ではない。フィールド `db: tx` は `*gorm.DB` をそのまま入れる。

`txRepo` は型の引数名。実行時は service 側の `repo` と同じ物。

## 5. 公式との対応

公式（Traditional API）は一塊の中に Create と `return` が全部ある。

```go
db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&Giraffe).Error; err != nil {
        return err
    }
    if err := tx.Create(&Lion).Error; err != nil {
        return err
    }
    return nil
})
```

007 は Create と for を `fn` に出し、repo は合図だけ転送する。

```go
return r.db.Transaction(func(tx *gorm.DB) error {
    return fn(&UserRepository{db: tx})
})
```

| 公式 | 007 |
|---|---|
| `tx.Create(&Giraffe)` / `Lion` | for の中の `repo.Create` |
| `return err` | Unmarshal / Create / `scanner.Err()` の失敗 |
| `return nil` | `fn` が `nil` を返す（=`return fn(...)` が `nil`） |

`return fn(...)` は「いつも成功」ではない。`fn` の戻りを GORM のコールバックの戻りにする。`fn` が `nil` なら公式の `return nil` と同じ。

```go
fn(...)
return nil  // ダメ。失敗しても COMMIT になる
```

## 6. `fn` は何か

`s.repo.Transaction(この関数)` で渡した関数。渡した時点ではまだ走らない。`return fn(...)` で初めて実行する。

副作用: ファイルを1行ずつ読み、INSERT を `tx` に積む。

戻り `error`: 副作用が成功したかの報告。ログ用ではなく、塊を失敗／成功にする信号。

- `nil` → 失敗なし → COMMIT
- `nil` でない → ROLLBACK（積んだ行も消える）

## 7. Transaction の動作フロー

```text
s.repo.Transaction(fn)          // fn はまだ動かない
  r.db.Transaction(...)         // GORM: BEGIN → tx
    fn(&UserRepository{db: tx}) // 下書き repo で for が走る
      return err または nil
  err → ROLLBACK / nil → COMMIT
  同じ値が Import の戻りになる
```

成功（5行とも OK）:

1. BEGIN
2. 5件を下書きに積む（まだ表には見えない）
3. `fn` が `nil`
4. COMMIT → 5件見える

失敗（例: 3行目で Unmarshal または Create が `error`）:

1. BEGIN
2. 1・2件目を下書き
3. `fn` が `return err`
4. ROLLBACK → 1・2件目も残らない

## 8. 関連コード

- ラップ: `internal/repository/user_repo.go` の `Transaction`
- 塊の中身: `internal/service/import_service.go` の `Import`（`os.Open` の外、for は中）
