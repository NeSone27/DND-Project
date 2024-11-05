package repositories

import (
	"database/sql"
	"todolist-service/models"
)

type QuestRepository struct {
	DB *sql.DB
}

func (r *QuestRepository) CreateQuest(quest *models.Quest) error {
	query := "INSERT INTO quest (title, description, difficulty_level_id, status, is_public) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	return r.DB.QueryRow(query, quest.Title, quest.Description, quest.DifficultyLevelID, quest.Status, quest.IsPublic).Scan(&quest.ID)
}

func (r *QuestRepository) GetQuests(isPublic bool) ([]models.Quest, error) {
	var query string
	if isPublic {
		query = "SELECT id, title, description, difficulty_level_id, status, is_public FROM quest WHERE status != 'archived' AND is_public = true"
	} else {
		query = "SELECT id, title, description, difficulty_level_id, status, is_public FROM quest WHERE status != 'archived'"
	}
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quests []models.Quest
	for rows.Next() {
		var quest models.Quest
		if err := rows.Scan(&quest.ID, &quest.Title, &quest.Description, &quest.DifficultyLevelID, &quest.Status, &quest.IsPublic); err != nil {
			return nil, err
		}
		quests = append(quests, quest)
	}
	return quests, nil
}

func (r *QuestRepository) GetQuestByID(id int) (*models.Quest, error) {
	var quest models.Quest
	query := "SELECT id, title, description, difficulty_level_id, status, is_public FROM quest WHERE id = $1 and status != 'archived'"
	err := r.DB.QueryRow(query, id).Scan(&quest.ID, &quest.Title, &quest.Description, &quest.DifficultyLevelID, &quest.Status, &quest.IsPublic)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &quest, nil
}

func (r *QuestRepository) UpdateQuest(quest *models.Quest) error {
	query := "UPDATE quest SET title = $1, description = $2, difficulty_level_id = $3, status = $4, is_public = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $6"
	_, err := r.DB.Exec(query, quest.Title, quest.Description, quest.DifficultyLevelID, quest.Status, quest.IsPublic, quest.ID)
	return err
}

func (r *QuestRepository) DeleteQuest(id int) error {
	query := "UPDATE quest SET status = 'archived', deleted_at = CURRENT_TIMESTAMP WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
