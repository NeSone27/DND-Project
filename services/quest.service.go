package services

import (
	"errors"
	"todolist-service/models"
	"todolist-service/repositories"
)

type QuestService struct {
	Repo      *repositories.QuestRepository
	RepoUser  *repositories.UserRepository
	RepoImage *repositories.ImageRepository
}

func (s *QuestService) CreateQuest(quest *models.CreateQuestRequest) error {
	if quest.Title == "" {
		return errors.New("title is required")
	}
	if quest.Status != "active" && quest.Status != "inactive" {
		return errors.New("status must be active or inactive")
	}
	if quest.DifficultyLevelID == 0 {
		return errors.New("difficulty_level_id is required")
	}
	if len(quest.Image) > 10 {
		return errors.New("image must be less than 10")
	}

	questToSave := &models.Quest{
		Title:             quest.Title,
		Description:       quest.Description,
		DifficultyLevelID: quest.DifficultyLevelID,
		Status:            quest.Status,
		IsPublic:          quest.IsPublic,
	}
	err := s.Repo.CreateQuest(questToSave)
	if err != nil {
		return err
	}
	if len(quest.Image) > 0 {
		for _, image := range quest.Image {
			imageToSave := &models.Image{
				QuestID: questToSave.ID,
				URL:     image,
			}
			err := s.RepoImage.CreateImage(imageToSave)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *QuestService) GetQuests(request models.GetQuestRequest) ([]models.Quest, error) {
	var isPublic bool
	if request.UserDNDID == nil {
		isPublic = true
	} else {
		isPublic = false
	}

	quests, err := s.Repo.GetQuests(isPublic)
	if err != nil {
		return nil, err
	}
	for i := range quests {
		images, err := s.RepoImage.GetImagesByQuestID(quests[i].ID)
		if err != nil {
			return nil, err
		}
		quests[i].Image = images
	}
	return quests, nil
}

func (s *QuestService) GetQuestByID(id int) (*models.Quest, error) {
	return s.Repo.GetQuestByID(id)
}

func (s *QuestService) UpdateQuest(quest *models.Quest) error {
	if len(quest.Image) > 10 {
		return errors.New("image must be less than 10")
	}
	if quest.Status != "active" && quest.Status != "inactive" {
		return errors.New("status must be active or inactive")
	}
	questToSave := &models.Quest{
		Title:             quest.Title,
		Description:       quest.Description,
		DifficultyLevelID: quest.DifficultyLevelID,
		Status:            quest.Status,
		IsPublic:          quest.IsPublic,
	}
	err := s.Repo.UpdateQuest(questToSave)
	if err != nil {
		return err
	}
	s.RepoImage.DeleteImageByQuestID(quest.ID)
	if len(quest.Image) > 0 {
		for _, image := range quest.Image {
			imageToSave := &models.Image{
				QuestID: quest.ID,
				URL:     image,
			}
			err := s.RepoImage.CreateImage(imageToSave)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *QuestService) DeleteQuest(id int, user *models.UserDND) error {
	quest, err := s.Repo.GetQuestByID(id)
	if err != nil || quest == nil {
		return errors.New("quest not found")
	}
	if quest.CreatedBy != user.ID && user.Role != "admin" {
		return errors.New("user is not authorized to delete this quest")
	}

	return s.Repo.DeleteQuest(id)
}
