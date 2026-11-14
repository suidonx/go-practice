package model

// Automigrationを行うために、Gormのモデルを定義する。
type User struct {
	ID   uint   `gorm:"primaryKey" json:"-"`
	Age  int    `json:"age"`
	Name string `gorm:"type:varchar(500)" json:"name"`
	Role string `gorm:"type:char(15)" json:"role"`
}

func (User) TableName() string {
	return "users"
}
