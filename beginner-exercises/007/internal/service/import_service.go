package service

import (
	"bufio"
	"encoding/json"
	"os"
	"github.com/suidonx/go-practice/beginner-exercises/007/internal/model"
	"github.com/suidonx/go-practice/beginner-exercises/007/internal/repository"
)

// サービスはユースケースを関数に記述する
// ユースケースはビジネスロジックを記述する
type ImportService struct {
	repo *repository.UserRepository
}

func NewImportService(repo *repository.UserRepository) *ImportService {
	return &ImportService{repo: repo}
}

func (s *ImportService) Import(path string) error {
	// ファイルを開く
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// ファイルを読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row model.LogRow
		if err := json.Unmarshal(scanner.Bytes(), &row); err != nil {
			return err
		}
		// ユーザーを作成する。
		if err := s.repo.Create(row.User); err != nil {
			return err
		}
	}
	// ファイルを閉じる
	// 正常終了ならscanner.Err()はnilを返す。読み取り失敗ならエラーを返す。
	return scanner.Err()
}
