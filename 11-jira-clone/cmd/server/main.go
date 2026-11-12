package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"11-jira-clone/internal/model"
	"11-jira-clone/internal/config"
)

func main() {

	cfg := config.Load() // config.go からDBの設定を読み込む
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName) // データベースの接続情報を生成
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // データベースに接続
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Project{}, &model.Task{}); err != nil { // データベースのマイグレーションを実行
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("database migrated successfully")

	// TODO: ルーティングを設定
	_ = db
}