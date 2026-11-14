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
