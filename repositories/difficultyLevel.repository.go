package repositories

import (
	"database/sql"
	"todolist-service/models"
)

type DifficultyLevelRepository struct {
	DB *sql.DB
}

func (r *DifficultyLevelRepository) CreateDifficultyLevel(difficultyLevel *models.DifficultyLevel) error {
	query := "INSERT INTO difficulty_level (name, detail, status) VALUES ($1, $2, $3) RETURNING id"
	return r.DB.QueryRow(query, difficultyLevel.Name, difficultyLevel.Detail, difficultyLevel.Status).Scan(&difficultyLevel.ID)
}

func (r *DifficultyLevelRepository) GetDifficultyLevels() ([]models.DifficultyLevel, error) {
	rows, err := r.DB.Query("SELECT id, name, detail, status FROM difficulty_level WHERE status != 'archived'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var difficultyLevels []models.DifficultyLevel
	for rows.Next() {
		var difficultyLevel models.DifficultyLevel
		if err := rows.Scan(&difficultyLevel.ID, &difficultyLevel.Name, &difficultyLevel.Detail, &difficultyLevel.Status); err != nil {
			return nil, err
		}
		difficultyLevels = append(difficultyLevels, difficultyLevel)
	}
	return difficultyLevels, nil
}

func (r *DifficultyLevelRepository) GetDifficultyLevelByID(id int) (*models.DifficultyLevel, error) {
	var difficultyLevel models.DifficultyLevel
	query := "SELECT id, name, detail, status FROM difficulty_level WHERE id = $1 and status != 'archived'"
	err := r.DB.QueryRow(query, id).Scan(&difficultyLevel.ID, &difficultyLevel.Name, &difficultyLevel.Detail, &difficultyLevel.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &difficultyLevel, nil
}

func (r *DifficultyLevelRepository) UpdateDifficultyLevel(difficultyLevel *models.DifficultyLevel) error {
	query := "UPDATE difficulty_level SET name = $1, detail = $2, status = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4"
	_, err := r.DB.Exec(query, difficultyLevel.Name, difficultyLevel.Detail, difficultyLevel.Status, difficultyLevel.ID)
	return err
}

func (r *DifficultyLevelRepository) DeleteDifficultyLevel(id int) error {
	query := "UPDATE difficulty_level SET status = 'archived', deleted_at = CURRENT_TIMESTAMP WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
