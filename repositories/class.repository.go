package repositories

import (
	"database/sql"
	"todolist-service/models"
)

type ClassRepository struct {
	DB *sql.DB
}

func (r *ClassRepository) CreateClass(class *models.Class) error {
	query := "INSERT INTO class (name, detail, status) VALUES ($1, $2, $3) RETURNING id"
	return r.DB.QueryRow(query, class.Name, class.Detail, class.Status).Scan(&class.ID)
}

func (r *ClassRepository) GetClasses() ([]models.Class, error) {
	rows, err := r.DB.Query("SELECT id, name, detail, status FROM class WHERE status != 'archived'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []models.Class
	for rows.Next() {
		var class models.Class
		if err := rows.Scan(&class.ID, &class.Name, &class.Detail, &class.Status); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}
	return classes, nil
}

func (r *ClassRepository) GetClassByID(id int) (*models.Class, error) {
	var class models.Class
	query := "SELECT id, name, detail, status FROM class WHERE id = $1 and status != 'archived'"
	err := r.DB.QueryRow(query, id).Scan(&class.ID, &class.Name, &class.Detail, &class.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *ClassRepository) UpdateClass(class *models.Class) error {
	query := "UPDATE class SET name = $1, detail = $2, status = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4"
	_, err := r.DB.Exec(query, class.Name, class.Detail, class.Status, class.ID)
	return err
}

func (r *ClassRepository) DeleteClass(id int) error {
	query := "UPDATE class SET status = 'archived', deleted_at = CURRENT_TIMESTAMP WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
