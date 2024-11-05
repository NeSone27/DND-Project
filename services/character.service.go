package services

import (
	"errors"
	"todolist-service/models"
	"todolist-service/repositories"
)

type CharacterService struct {
	Repo      *repositories.CharacterRepository
	RepoUser  *repositories.UserRepository
	RepoImage *repositories.ImageRepository
}

func (s *CharacterService) CreateCharacter(character *models.CreateUserCharacterRequest) error {
	if character.Title == "" {
		return errors.New("title is required")
	}
	if character.Status != "active" && character.Status != "inactive" {
		return errors.New("status must be active or inactive")
	}
	if character.ClassID == 0 {
		return errors.New("class_id is required")
	}
	if character.RaceID == 0 {
		return errors.New("race_id is required")
	}
	if len(character.Image) > 10 {
		return errors.New("image must be less than 10")
	}

	characterToSave := &models.UserCharacter{
		Title:       character.Title,
		Description: character.Description,
		ClassID:     character.ClassID,
		RaceID:      character.RaceID,
		Status:      character.Status,
		IsPublic:    character.IsPublic,
		CreatedBy:   character.CreatedBy,
	}
	err := s.Repo.CreateCharacter(characterToSave)
	if err != nil {
		return err
	}
	if len(character.Image) > 0 {
		for _, image := range character.Image {
			imageToSave := &models.Image{
				CharacterID: characterToSave.ID,
				URL:         image,
			}
			err := s.RepoImage.CreateImage(imageToSave)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *CharacterService) GetCharacters(request models.GetUserCharacterRequest) ([]models.UserCharacter, error) {
	var isPublic bool
	if request.UserDNDID == nil {
		isPublic = true
	} else {
		isPublic = false
	}

	characters, err := s.Repo.GetCharacters(isPublic)
	if err != nil {
		return nil, err
	}
	for i := range characters {
		images, err := s.RepoImage.GetImagesByCharacterID(characters[i].ID)
		if err != nil {
			return nil, err
		}
		characters[i].Image = images
	}
	return characters, nil
}

func (s *CharacterService) GetCharacterByID(id int) (*models.UserCharacter, error) {
	return s.Repo.GetCharacterByID(id)
}

func (s *CharacterService) UpdateCharacter(character *models.UserCharacter) error {
	if len(character.Image) > 10 {
		return errors.New("image must be less than 10")
	}
	if character.Status != "active" && character.Status != "inactive" {
		return errors.New("status must be active or inactive")
	}
	characterToSave := &models.UserCharacter{
		Title:       character.Title,
		Description: character.Description,
		ClassID:     character.ClassID,
		RaceID:      character.RaceID,
		Status:      character.Status,
		IsPublic:    character.IsPublic,
	}
	err := s.Repo.UpdateCharacter(characterToSave)
	if err != nil {
		return err
	}
	s.RepoImage.DeleteImageByCharacterID(character.ID)
	if len(character.Image) > 0 {
		for _, image := range character.Image {
			imageToSave := &models.Image{
				CharacterID: character.ID,
				URL:         image,
			}
			err := s.RepoImage.CreateImage(imageToSave)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *CharacterService) DeleteCharacter(id int, user *models.UserDND) error {
	character, err := s.Repo.GetCharacterByID(id)
	if err != nil || character == nil {
		return errors.New("character not found")
	}
	if character.CreatedBy != user.ID && user.Role != "admin" {
		return errors.New("user is not authorized to delete this character")
	}

	return s.Repo.DeleteCharacter(id)
}
