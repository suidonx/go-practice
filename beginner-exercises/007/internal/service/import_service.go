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

	return s.repo.Transaction(func(repo *repository.UserRepository) error {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var row model.LogRow
			if err := json.Unmarshal(scanner.Bytes(), &row); err != nil {
				return err
			}
			if err := repo.Create(row.User); err != nil {
				return err
			}
		}
		return scanner.Err()
	})
}
