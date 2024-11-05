package repositories

import (
	"database/sql"
	"todolist-service/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *models.UserDND) error {
	query := "INSERT INTO user_dnd (username, password, role, status) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.DB.QueryRow(query, user.Username, user.Password, user.Role, user.Status).Scan(&user.ID)
}

func (r *UserRepository) GetUsers() ([]models.UserDND, error) {
	rows, err := r.DB.Query("SELECT id, username, password, role, status FROM user_dnd WHERE status != 'archived'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserDND
	for rows.Next() {
		var user models.UserDND
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Status); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.UserDND, error) {
	var user models.UserDND
	query := "SELECT id, username, password, role, status FROM user_dnd WHERE id = $1 and status != 'archived'"
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.UserDND) error {
	query := "UPDATE user_dnd SET username = $1, password = $2, role = $3, status = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5"
	_, err := r.DB.Exec(query, user.Username, user.Password, user.Role, user.Status, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	query := "UPDATE user_dnd SET status = 'archived', deleted_at = CURRENT_TIMESTAMP WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
