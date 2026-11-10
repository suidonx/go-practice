# Contributing

## フォルダ構成

- 章: `01-basics-functions` のように **番号 + ハイフン + トピック名**
- 章の中: `package` / `pointers` など **トピック名のみ**（番号なし）
- **各トピックフォルダに `main.go` を置く**（`package main`）。新規トピックを作るときはフォルダと同時に追加する
- メモ: トピック名に合わせた `.md`（例: `zeros/zero.md`）は任意

## コミットメッセージ（Conventional Commits）

```
<type>(<scope>): <subject>
```

### type の例

| type | 使うとき |
|------|----------|
| `feat` | 新しい練習トピック・`main.go` の追加 |
| `docs` | `.md` や README のみの更新 |
| `refactor` | フォルダ名・構成の整理（内容は同じ） |
| `chore` | リポジトリまわりの雑務 |

### 例

```
feat(variables): add constants exercise
docs(zeros): explain zero value benefits
refactor: rename chapter dirs to hyphen style
```

- `subject` は英語・簡潔に（命令形が一般的）
- 1 コミット 1 意図
