package repositories

import (
	"database/sql"
	"todolist-service/models"
)

type RaceRepository struct {
	DB *sql.DB
}

func (r *RaceRepository) CreateRace(race *models.Race) error {
	query := "INSERT INTO race (name, detail, status) VALUES ($1, $2, $3) RETURNING id"
	return r.DB.QueryRow(query, race.Name, race.Detail, race.Status).Scan(&race.ID)
}

func (r *RaceRepository) GetRaces() ([]models.Race, error) {
	rows, err := r.DB.Query("SELECT id, name, detail, status FROM race WHERE status != 'archived'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var races []models.Race
	for rows.Next() {
		var race models.Race
		if err := rows.Scan(&race.ID, &race.Name, &race.Detail, &race.Status); err != nil {
			return nil, err
		}
		races = append(races, race)
	}
	return races, nil
}

func (r *RaceRepository) GetRaceByID(id int) (*models.Race, error) {
	var race models.Race
	query := "SELECT id, name, detail, status FROM race WHERE id = $1 and status != 'archived'"
	err := r.DB.QueryRow(query, id).Scan(&race.ID, &race.Name, &race.Detail, &race.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &race, nil
}

func (r *RaceRepository) UpdateRace(race *models.Race) error {
	query := "UPDATE race SET name = $1, detail = $2, status = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4"
	_, err := r.DB.Exec(query, race.Name, race.Detail, race.Status, race.ID)
	return err
}

func (r *RaceRepository) DeleteRace(id int) error {
	query := "UPDATE race SET status = 'archived', deleted_at = CURRENT_TIMESTAMP WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
