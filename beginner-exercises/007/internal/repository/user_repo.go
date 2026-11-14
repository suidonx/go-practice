package repository

import (
	"github.com/suidonx/go-practice/beginner-exercises/007/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// ユーザーリポジトリを初期化する。
// コンストラクタ
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// ユーザーテーブルを作成する。
func (r *UserRepository) AutoMigrate() error {
	return r.db.AutoMigrate(&model.User{})
}

// ユーザーを作成する。
func (r *UserRepository) Create(u model.User) error {
	return r.db.Create(&u).Error
}

// db.Createと同様にdb.Transactionをラップする。
// why: サービス層にGORMを見せないため。repo.Transactionを呼べるようにして
// サービス層ではトランザクションを意識する必要がなくなる。

/*
ラップされているが、実体はreturn fn(&UserRepository{db: tx})を呼び出している。
Service層のメソッドをfnに渡している。
結局、Service層のメソッドの処理がトランザクション内で実行される。

最後のreturnでエラーが返った場合はrollback、nilが返った場合はcommitされる。
*/

// Transaction は fn を1つのトランザクションで実行する。
// fn が error を返すか panic すると rollback、nil なら commit。
func (r *UserRepository) Transaction(fn func(txRepo *UserRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(&UserRepository{db: tx})

	})
}
