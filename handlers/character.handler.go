package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist-service/models"
	"todolist-service/services"

	"github.com/gorilla/mux"
)

type CharacterHandler struct {
	Service     *services.CharacterService
	ServiceUser *services.UserService
}

func NewCharacterHandler(characterService *services.CharacterService, userService *services.UserService) *CharacterHandler {
	return &CharacterHandler{
		Service:     characterService,
		ServiceUser: userService,
	}
}

func (h *CharacterHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var character models.CreateUserCharacterRequest
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if character.CreatedBy == 0 {
		http.Error(w, "created_by is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(character.CreatedBy)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" && user.Role != "user" {
		http.Error(w, "user is not registered", http.StatusForbidden)
		return
	}
	characterToSave := &models.CreateUserCharacterRequest{
		Title:       character.Title,
		Description: character.Description,
		ClassID:     character.ClassID,
		RaceID:      character.RaceID,
		Status:      character.Status,
		IsPublic:    character.IsPublic,
		Image:       character.Image,
		CreatedBy:   character.CreatedBy,
	}
	if err := h.Service.CreateCharacter(characterToSave); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(characterToSave)
}

func (h *CharacterHandler) GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	character, err := h.Service.GetCharacterByID(id)
	if err != nil || character == nil {
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(character)
}

func (h *CharacterHandler) GetCharacters(w http.ResponseWriter, r *http.Request) {
	var request models.GetUserCharacterRequest
	if r.URL.Query().Get("user_dnd_id") != "" {
		userDNDID, err := strconv.Atoi(r.URL.Query().Get("user_dnd_id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		request.UserDNDID = &userDNDID
	}
	characters, err := h.Service.GetCharacters(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}

func (h *CharacterHandler) UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var character models.UpdateUserCharacterRequest
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if character.UpdatedBy == 0 {
		http.Error(w, "updated_by is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(character.UpdatedBy)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" && user.Role != "user" {
		http.Error(w, "user is not registered", http.StatusForbidden)
		return
	}

	existingCharacter, err := h.Service.GetCharacterByID(id)
	if err != nil || existingCharacter == nil {
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}
	if existingCharacter.CreatedBy != character.UpdatedBy && user.Role != "admin" {
		http.Error(w, "user is not authorized to update this character", http.StatusForbidden)
		return
	}

	existingCharacter.ID = id
	if character.Title != "" {
		existingCharacter.Title = character.Title
	}
	if character.Description != "" {
		existingCharacter.Description = character.Description
	}
	if character.Status != "" {
		existingCharacter.Status = character.Status
	}
	if character.IsPublic != existingCharacter.IsPublic {
		existingCharacter.IsPublic = character.IsPublic
	}
	if character.ClassID != 0 {
		existingCharacter.ClassID = character.ClassID
	}
	if character.RaceID != 0 {
		existingCharacter.RaceID = character.RaceID
	}
	existingCharacter.Image = character.Image
	if err := h.Service.UpdateCharacter(existingCharacter); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingCharacter)
}

func (h *CharacterHandler) DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if r.Body == nil || r.Body == http.NoBody {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	var character models.GetUserCharacterRequest
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if character.UserDNDID == nil {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(*character.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	err = h.Service.DeleteCharacter(id, user)
	if err != nil {
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Character deleted successfully"})
}
