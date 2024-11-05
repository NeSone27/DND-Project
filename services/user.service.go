package services

import (
	"errors"
	"strings"
	"todolist-service/models"
	"todolist-service/repositories"
	"todolist-service/utils"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) CreateUser(user *models.UserDND) error {
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Role == "" {
		return errors.New("role is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if strings.Contains(user.Username, " ") {
		return errors.New("username cannot contain spaces")
	}
	if user.Role != "admin" && user.Role != "user" {
		return errors.New("role must be admin or user")
	}
	if user.Status != "active" && user.Status != "inactive" {
		return errors.New("status must be active or inactive")
	}
	user.Password = utils.HashPassword(user.Password)

	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUsers() ([]models.UserDND, error) {
	return s.Repo.GetUsers()
}

func (s *UserService) GetUserByID(id int) (*models.UserDND, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *models.UserDND) error {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	user, err := s.Repo.GetUserByID(id)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	return s.Repo.DeleteUser(id)
}
