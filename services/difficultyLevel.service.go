package services

import (
	"errors"
	"todolist-service/models"
	"todolist-service/repositories"
)

type DifficultyLevelService struct {
	Repo     *repositories.DifficultyLevelRepository
	RepoUser *repositories.UserRepository
}

func (s *DifficultyLevelService) CreateDifficultyLevel(difficultyLevel *models.DifficultyLevel) error {
	if difficultyLevel.Name == "" {
		return errors.New("name is required")
	}
	if difficultyLevel.Status != "active" && difficultyLevel.Status != "inactive" {
		return errors.New("status must be active or inactive")
	}
	difficultyLevelToSave := &models.DifficultyLevel{
		Name:   difficultyLevel.Name,
		Detail: difficultyLevel.Detail,
		Status: difficultyLevel.Status,
	}

	return s.Repo.CreateDifficultyLevel(difficultyLevelToSave)
}

func (s *DifficultyLevelService) GetDifficultyLevels() ([]models.DifficultyLevel, error) {
	return s.Repo.GetDifficultyLevels()
}

func (s *DifficultyLevelService) GetDifficultyLevelByID(id int) (*models.DifficultyLevel, error) {
	return s.Repo.GetDifficultyLevelByID(id)
}

func (s *DifficultyLevelService) UpdateDifficultyLevel(difficultyLevel *models.DifficultyLevel) error {
	return s.Repo.UpdateDifficultyLevel(difficultyLevel)
}

func (s *DifficultyLevelService) DeleteDifficultyLevel(id int) error {
	difficultyLevel, err := s.Repo.GetDifficultyLevelByID(id)
	if err != nil || difficultyLevel == nil {
		return errors.New("difficulty level not found")
	}

	return s.Repo.DeleteDifficultyLevel(id)
}
