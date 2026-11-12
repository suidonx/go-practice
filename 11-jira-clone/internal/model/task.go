package model

import "time"

type Task struct {
	ID uint `gorm:"primaryKey" json: "id"`
	ProjectID uint `gorm:"not null" json: "project_id"`
	Title string `gorm:"not null" json: "title"`
	Description string `gorm:"type:text" json: "description"`
	Status string `gorm:"not null" json: "status"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

func (Task) TableName() string {
	return "tasks"
}