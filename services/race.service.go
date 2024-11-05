package services

import (
	"errors"
	"todolist-service/models"
	"todolist-service/repositories"
)

type RaceService struct {
	Repo     *repositories.RaceRepository
	RepoUser *repositories.UserRepository
}

func (s *RaceService) CreateRace(race *models.Race) error {
	if race.Name == "" {
		return errors.New("name is required")
	}
	if race.Status != "active" && race.Status != "inactive" {
		return errors.New("status must be active or inactive")
	}
	raceToSave := &models.Race{
		Name:   race.Name,
		Detail: race.Detail,
		Status: race.Status,
	}

	return s.Repo.CreateRace(raceToSave)
}

func (s *RaceService) GetRaces() ([]models.Race, error) {
	return s.Repo.GetRaces()
}

func (s *RaceService) GetRaceByID(id int) (*models.Race, error) {
	return s.Repo.GetRaceByID(id)
}

func (s *RaceService) UpdateRace(race *models.Race) error {
	return s.Repo.UpdateRace(race)
}

func (s *RaceService) DeleteRace(id int) error {
	race, err := s.Repo.GetRaceByID(id)
	if err != nil || race == nil {
		return errors.New("race not found")
	}

	return s.Repo.DeleteRace(id)
}
