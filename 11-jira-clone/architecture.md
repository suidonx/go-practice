# 構成

ディレクトリ構成と、それが何のアーキテクチャかをまとめる。

## ディレクトリ

```
.
├── cmd
│   └── server
│       └── main.go              # エントリポイント（サーバー起動）
├── config
│   └── config.go                # 環境変数・DB設定
├── internal
│   ├── controller
│   │   ├── project_handler.go   # HTTPリクエスト制御
│   │   └── task_handler.go
│   ├── model
│   │   ├── project.go           # GORM構造体定義
│   │   └── task.go
│   ├── repository
│   │   ├── project_repo.go      # DB操作（SQL発行）
│   │   └── task_repo.go
│   ├── router
│   │   └── router.go            # ルーティング定義（Gin）
│   └── service
│       ├── project_service.go   # ビジネスロジック
│       └── task_service.go
├── docker-compose.yml           # PostgreSQL環境
├── go.mod                       # 依存管理
└── go.sum
```

各層の役割:

- `controller` … HTTP の入り口
- `service` … 業務の処理
- `repository` … DB アクセス
- `model` … テーブルに近い構造体（GORM）

リクエストの流れ:

```
router → controller → service → repository → DB
```

Go の Web API でよく見る分け方である。HTTP と DB を同じファイルに混ぜない、という点がレイヤー分けの入口になる。
