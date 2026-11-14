package model

// ログデータを保存するためのモデルを定義する。
type LogRow struct {
	User User `json:"user"`
}
