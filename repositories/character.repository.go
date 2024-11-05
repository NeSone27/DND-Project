package repositories

import (
	"database/sql"
	"todolist-service/models"
)

type CharacterRepository struct {
	DB *sql.DB
}

func (r *CharacterRepository) CreateCharacter(character *models.UserCharacter) error {
	query := "INSERT INTO character (title, description, class_id, race_id, status, is_public) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	return r.DB.QueryRow(query, character.Title, character.Description, character.ClassID, character.RaceID, character.Status, character.IsPublic).Scan(&character.ID)
}

func (r *CharacterRepository) GetCharacters(isPublic bool) ([]models.UserCharacter, error) {
	var query string
	if isPublic {
		query = "SELECT id, title, description, class_id, race_id, status, is_public FROM character WHERE status != 'archived' AND is_public = true"
	} else {
		query = "SELECT id, title, description, class_id, race_id, status, is_public FROM character WHERE status != 'archived'"
	}
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []models.UserCharacter
	for rows.Next() {
		var character models.UserCharacter
		if err := rows.Scan(&character.ID, &character.Title, &character.Description, &character.ClassID, &character.RaceID, &character.Status, &character.IsPublic); err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}
	return characters, nil
}

func (r *CharacterRepository) GetCharacterByID(id int) (*models.UserCharacter, error) {
	var character models.UserCharacter
	query := "SELECT id, title, description, class_id, race_id, status, is_public FROM character WHERE id = $1 and status != 'archived'"
	err := r.DB.QueryRow(query, id).Scan(&character.ID, &character.Title, &character.Description, &character.ClassID, &character.RaceID, &character.Status, &character.IsPublic)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func (r *CharacterRepository) UpdateCharacter(character *models.UserCharacter) error {
	query := "UPDATE character SET title = $1, description = $2, class_id = $3, race_id = $4, status = $5, is_public = $6, updated_at = CURRENT_TIMESTAMP WHERE id = $7"
	_, err := r.DB.Exec(query, character.Title, character.Description, character.ClassID, character.RaceID, character.Status, character.IsPublic, character.ID)
	return err
}

func (r *CharacterRepository) DeleteCharacter(id int) error {
	query := "UPDATE character SET status = 'archived', deleted_at = CURRENT_TIMESTAMP WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
