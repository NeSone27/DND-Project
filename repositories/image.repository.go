package repositories

import (
	"database/sql"
	"todolist-service/models"
)

type ImageRepository struct {
	DB *sql.DB
}

func (r *ImageRepository) CreateImage(image *models.Image) error {
	query := "INSERT INTO image (character_id, quest_id, url) VALUES ($1, $2, $3) RETURNING id"
	return r.DB.QueryRow(query, image.CharacterID, image.QuestID, image.URL).Scan(&image.ID)
}

func (r *ImageRepository) GetImages() ([]models.Image, error) {
	rows, err := r.DB.Query("SELECT id, character_id, quest_id, url FROM image")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []models.Image
	for rows.Next() {
		var image models.Image
		if err := rows.Scan(&image.ID, &image.CharacterID, &image.QuestID, &image.URL); err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}

func (r *ImageRepository) GetImageByID(id int) (*models.Image, error) {
	var image models.Image
	query := "SELECT id, character_id, quest_id, url FROM image WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&image.ID, &image.CharacterID, &image.QuestID, &image.URL)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &image, nil
}

func (r *ImageRepository) UpdateImage(image *models.Image) error {
	query := "UPDATE image SET character_id = $1, quest_id = $2, url = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4"
	_, err := r.DB.Exec(query, image.CharacterID, image.QuestID, image.URL, image.ID)
	return err
}

func (r *ImageRepository) DeleteImage(id int) error {
	query := "UPDATE image SET status = 'archived', deleted_at = CURRENT_TIMESTAMP WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *ImageRepository) DeleteImageByCharacterID(id int) error {
	query := "DELETE FROM image WHERE character_id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *ImageRepository) GetImagesByCharacterID(id int) ([]string, error) {
	query := "SELECT url FROM image WHERE character_id = $1"
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []string
	for rows.Next() {
		var image string
		if err := rows.Scan(&image); err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}

func (r *ImageRepository) GetImagesByQuestID(id int) ([]string, error) {
	query := "SELECT url FROM image WHERE quest_id = $1"
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []string
	for rows.Next() {
		var image string
		if err := rows.Scan(&image); err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}

func (r *ImageRepository) DeleteImageByQuestID(id int) error {
	query := "DELETE FROM image WHERE quest_id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
