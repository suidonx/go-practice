package model

import "time"

type Project struct {
	ID uint `gorm:"primaryKey" json: "id"` 
	Name string `gorm:"not null" json: "name"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

// why: テーブル名を指定するためのメソッド
// gormではデフォルトで構造体名の複数形(Projects)をテーブル名として使用するが、これを上書きしている
func (Project) TableName() string {
	return "projects"
}