package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist-service/models"
	"todolist-service/services"

	"github.com/gorilla/mux"
)

type QuestHandler struct {
	Service     *services.QuestService
	ServiceUser *services.UserService
}

func NewQuestHandler(questService *services.QuestService, userService *services.UserService) *QuestHandler {
	return &QuestHandler{
		Service:     questService,
		ServiceUser: userService,
	}
}

func (h *QuestHandler) CreateQuest(w http.ResponseWriter, r *http.Request) {
	var quest models.CreateQuestRequest
	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if quest.CreatedBy == 0 {
		http.Error(w, "created_by is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(quest.CreatedBy)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" && user.Role != "user" {
		http.Error(w, "user is not registered", http.StatusForbidden)
		return
	}
	questToSave := &models.CreateQuestRequest{
		Title:             quest.Title,
		Description:       quest.Description,
		DifficultyLevelID: quest.DifficultyLevelID,
		Status:            quest.Status,
		IsPublic:          quest.IsPublic,
		Image:             quest.Image,
		CreatedBy:         quest.CreatedBy,
	}
	if err := h.Service.CreateQuest(questToSave); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(questToSave)
}

func (h *QuestHandler) GetQuestByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	quest, err := h.Service.GetQuestByID(id)
	if err != nil || quest == nil {
		http.Error(w, "Quest not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quest)
}

func (h *QuestHandler) GetQuests(w http.ResponseWriter, r *http.Request) {
	var request models.GetQuestRequest
	if r.URL.Query().Get("user_dnd_id") != "" {
		userDNDID, err := strconv.Atoi(r.URL.Query().Get("user_dnd_id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		request.UserDNDID = &userDNDID
	}
	quests, err := h.Service.GetQuests(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quests)
}

func (h *QuestHandler) UpdateQuest(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var quest models.UpdateQuestRequest
	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if quest.UpdatedBy == 0 {
		http.Error(w, "updated_by is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(quest.UpdatedBy)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" && user.Role != "user" {
		http.Error(w, "user is not registered", http.StatusForbidden)
		return
	}

	existingQuest, err := h.Service.GetQuestByID(id)
	if err != nil || existingQuest == nil {
		http.Error(w, "Quest not found", http.StatusNotFound)
		return
	}
	if existingQuest.CreatedBy != quest.UpdatedBy && user.Role != "admin" {
		http.Error(w, "user is not authorized to update this quest", http.StatusForbidden)
		return
	}

	existingQuest.ID = id
	if quest.Title != "" {
		existingQuest.Title = quest.Title
	}
	if quest.Description != "" {
		existingQuest.Description = quest.Description
	}
	if quest.Status != "" {
		existingQuest.Status = quest.Status
	}
	if quest.IsPublic != existingQuest.IsPublic {
		existingQuest.IsPublic = quest.IsPublic
	}
	if quest.DifficultyLevelID != 0 {
		existingQuest.DifficultyLevelID = quest.DifficultyLevelID
	}
	existingQuest.Image = quest.Image
	if err := h.Service.UpdateQuest(existingQuest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingQuest)
}

func (h *QuestHandler) DeleteQuest(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if r.Body == nil || r.Body == http.NoBody {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	var quest models.GetQuestRequest
	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if quest.UserDNDID == nil {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(*quest.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	err = h.Service.DeleteQuest(id, user)
	if err != nil {
		http.Error(w, "Quest not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Quest deleted successfully"})
}
