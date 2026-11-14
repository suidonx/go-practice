package main

import (
	"fmt"
	"github.com/suidonx/go-practice/beginner-exercises/007/config"
	"github.com/suidonx/go-practice/beginner-exercises/007/internal/repository"
	"github.com/suidonx/go-practice/beginner-exercises/007/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// コマンドラインオプションからファイルパスを取得
	if len(os.Args) < 2 {
		log.Fatal("ファイルパスを指定してください")
	}
	if len(os.Args) > 2 {
		log.Fatal("複数のオプションは指定できません")
	}
	filePath := os.Args[1]

	// DB接続
	cfg := config.Load()
	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("DB接続に失敗しました: ", err)
	}

	// DB接続を閉じる
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("DBクローズに失敗しました: ", err)
	}
	defer sqlDB.Close()

	// ユーザーレポジトリを初期化
	userRepository := repository.NewUserRepository(db)

	// AutoMigrate
	err = userRepository.AutoMigrate()
	if err != nil {
		log.Fatal("テーブル作成に失敗しました: ", err)
	}

	// importサービスを初期化
	importService := service.NewImportService(userRepository)

	// importサービスのインポートを呼ぶ
	if err := importService.Import(filePath); err != nil {
		// エラーが発生した場合はログを出力して終了
		log.Fatal("インポートに失敗しました: ", err)
	}

	// 正常終了時のメッセージ
	fmt.Println("インポートが完了しました")
}
