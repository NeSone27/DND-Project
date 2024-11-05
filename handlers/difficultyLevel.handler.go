package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist-service/models"
	"todolist-service/services"

	"github.com/gorilla/mux"
)

type DifficultyLevelHandler struct {
	Service     *services.DifficultyLevelService
	ServiceUser *services.UserService
}

func NewDifficultyLevelHandler(difficultyLevelService *services.DifficultyLevelService, userService *services.UserService) *DifficultyLevelHandler {
	return &DifficultyLevelHandler{
		Service:     difficultyLevelService,
		ServiceUser: userService,
	}
}

func (h *DifficultyLevelHandler) CreateDifficultyLevel(w http.ResponseWriter, r *http.Request) {
	var difficultyLevel models.DifficultyLevelRequest
	if err := json.NewDecoder(r.Body).Decode(&difficultyLevel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if difficultyLevel.UserDNDID == 0 {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(difficultyLevel.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "user is not admin", http.StatusForbidden)
		return
	}

	difficultyLevelToSave := &models.DifficultyLevel{
		Name:   difficultyLevel.Name,
		Detail: difficultyLevel.Detail,
		Status: difficultyLevel.Status,
	}
	if err := h.Service.CreateDifficultyLevel(difficultyLevelToSave); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(difficultyLevelToSave)
}

func (h *DifficultyLevelHandler) GetDifficultyLevelByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	difficultyLevel, err := h.Service.GetDifficultyLevelByID(id)
	if err != nil || difficultyLevel == nil {
		http.Error(w, "Difficulty level not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(difficultyLevel)
}

func (h *DifficultyLevelHandler) GetDifficultyLevels(w http.ResponseWriter, r *http.Request) {
	difficultyLevels, err := h.Service.GetDifficultyLevels()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(difficultyLevels)
}

func (h *DifficultyLevelHandler) UpdateDifficultyLevel(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var difficultyLevel models.DifficultyLevelRequest
	if err := json.NewDecoder(r.Body).Decode(&difficultyLevel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if difficultyLevel.UserDNDID == 0 {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(difficultyLevel.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "user is not admin", http.StatusForbidden)
		return
	}

	existingDifficultyLevel, err := h.Service.GetDifficultyLevelByID(id)
	if err != nil || existingDifficultyLevel == nil {
		http.Error(w, "Difficulty level not found", http.StatusNotFound)
		return
	}

	existingDifficultyLevel.ID = id
	if difficultyLevel.Name != "" {
		existingDifficultyLevel.Name = difficultyLevel.Name
	}
	if difficultyLevel.Detail != "" {
		existingDifficultyLevel.Detail = difficultyLevel.Detail
	}
	if difficultyLevel.Status != "" {
		existingDifficultyLevel.Status = difficultyLevel.Status
	}
	if err := h.Service.UpdateDifficultyLevel(existingDifficultyLevel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingDifficultyLevel)
}

func (h *DifficultyLevelHandler) DeleteDifficultyLevel(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if r.Body == nil || r.Body == http.NoBody {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	var difficultyLevel models.DifficultyLevelRequest
	if err := json.NewDecoder(r.Body).Decode(&difficultyLevel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if difficultyLevel.UserDNDID == 0 {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(difficultyLevel.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "user is not admin", http.StatusForbidden)
		return
	}

	err = h.Service.DeleteDifficultyLevel(id)
	if err != nil {
		http.Error(w, "Difficulty level not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Difficulty level deleted successfully"})
}
