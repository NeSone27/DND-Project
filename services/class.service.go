package services

import (
	"errors"
	"todolist-service/models"
	"todolist-service/repositories"
)

type ClassService struct {
	Repo     *repositories.ClassRepository
	RepoUser *repositories.UserRepository
}

func (s *ClassService) CreateClass(class *models.Class) error {
	if class.Name == "" {
		return errors.New("name is required")
	}
	if class.Status != "active" && class.Status != "inactive" {
		return errors.New("status must be active or inactive")
	}
	classToSave := &models.Class{
		Name:   class.Name,
		Detail: class.Detail,
		Status: class.Status,
	}

	return s.Repo.CreateClass(classToSave)
}

func (s *ClassService) GetClasses() ([]models.Class, error) {
	return s.Repo.GetClasses()
}

func (s *ClassService) GetClassByID(id int) (*models.Class, error) {
	return s.Repo.GetClassByID(id)
}

func (s *ClassService) UpdateClass(class *models.Class) error {
	return s.Repo.UpdateClass(class)
}

func (s *ClassService) DeleteClass(id int) error {
	class, err := s.Repo.GetClassByID(id)
	if err != nil || class == nil {
		return errors.New("class not found")
	}

	return s.Repo.DeleteClass(id)
}
