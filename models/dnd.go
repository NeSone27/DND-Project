package models

type RaceRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	Status    string `json:"status"`
	UserDNDID int    `json:"user_dnd_id"`
}
type Race struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

type ClassRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	Status    string `json:"status"`
	UserDNDID int    `json:"user_dnd_id"`
}
type Class struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

type DifficultyLevelRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	Status    string `json:"status"`
	UserDNDID int    `json:"user_dnd_id"`
}
type DifficultyLevel struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

type UserDND struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}
type CreateQuestRequest struct {
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	DifficultyLevelID int      `json:"difficulty_level_id"`
	Status            string   `json:"status"`
	IsPublic          bool     `json:"is_public"`
	CreatedBy         int      `json:"created_by"`
	Image             []string `json:"image"`
}
type UpdateQuestRequest struct {
	ID                int      `json:"id"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	DifficultyLevelID int      `json:"difficulty_level_id"`
	Status            string   `json:"status"`
	IsPublic          bool     `json:"is_public"`
	UpdatedBy         int      `json:"updated_by"`
	Image             []string `json:"image"`
}
type GetQuestRequest struct {
	UserDNDID *int `json:"user_dnd_id"`
}
type Quest struct {
	ID                int      `json:"id"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	DifficultyLevelID int      `json:"difficulty_level_id"`
	Status            string   `json:"status"`
	IsPublic          bool     `json:"is_public"`
	CreatedBy         int      `json:"created_by"`
	Image             []string `json:"image"`
}
type CreateUserCharacterRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ClassID     int      `json:"class_id"`
	RaceID      int      `json:"race_id"`
	Status      string   `json:"status"`
	IsPublic    bool     `json:"is_public"`
	CreatedBy   int      `json:"created_by"`
	Image       []string `json:"image"`
}

type UpdateUserCharacterRequest struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ClassID     int      `json:"class_id"`
	RaceID      int      `json:"race_id"`
	Status      string   `json:"status"`
	IsPublic    bool     `json:"is_public"`
	UpdatedBy   int      `json:"updated_by"`
	Image       []string `json:"image"`
}

type GetUserCharacterRequest struct {
	UserDNDID *int `json:"user_dnd_id"`
}

type UserCharacter struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ClassID     int      `json:"class_id"`
	RaceID      int      `json:"race_id"`
	Status      string   `json:"status"`
	IsPublic    bool     `json:"is_public"`
	CreatedBy   int      `json:"created_by"`
	Image       []string `json:"image"`
}

type Image struct {
	ID          int    `json:"id"`
	QuestID     int    `json:"quest_id"`
	CharacterID int    `json:"character_id"`
	URL         string `json:"url"`
}
