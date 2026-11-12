package repository

import (
	"gorm.io/gorm"
)

// ProjectRepositoryはプロジェクトのCRUDを行う
type ProjectRepository struct {
	db *gorm.DB // データベースの接続情報
}


// 外から db を受け取って、リポジトリを作る関数である。New で始まるのはコンストラクタ相当の慣習である。
func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	// 構造体を作り、そのポインタを返す。引数の db をフィールドに入れている。
	return &ProjectRepository{db: db}
}

// プロジェクトの一覧を取得する
func (r *ProjectRepository) List() ([]model.Project, error) {
	var projects []model.Project
	if err := r.db.Order("id").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// プロジェクトをIDで取得する
func (r *ProjectRepository) GetByID(id uint) (*model.Project, error) {
	var p model.Project
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
