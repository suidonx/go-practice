package repository

import (

)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) ListByProjectID(projectID uint) ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Where("project_id = ?", projectID).Order("id").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetByID(id uint) (*model.Task, error) {
	var t model.Task
	if err := r.db.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}