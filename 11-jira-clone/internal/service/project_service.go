package service

import (
	"11-jira-clone/internal/repository"
	"11-jira-clone/internal/model"
)

type ProjectService struct {
	projectRepo *repository.ProjectRepository
}

func NewProjectService(projectRepo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}

func (s *ProjectService) List() ([]model.Project, error) {
	return s.projectRepo.List()
}

func (s *ProjectService) GetByID(id uint) (*model.Project, error) {
	return s.projectRepo.GetByID(id)
}