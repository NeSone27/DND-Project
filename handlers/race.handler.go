package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist-service/models"
	"todolist-service/services"

	"github.com/gorilla/mux"
)

type RaceHandler struct {
	Service     *services.RaceService
	ServiceUser *services.UserService
}

func NewRaceHandler(raceService *services.RaceService, userService *services.UserService) *RaceHandler {
	return &RaceHandler{
		Service:     raceService,
		ServiceUser: userService,
	}
}

func (h *RaceHandler) CreateRace(w http.ResponseWriter, r *http.Request) {
	var race models.RaceRequest
	if err := json.NewDecoder(r.Body).Decode(&race); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if race.UserDNDID == 0 {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(race.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "user is not admin", http.StatusForbidden)
		return
	}

	raceToSave := &models.Race{
		Name:   race.Name,
		Detail: race.Detail,
		Status: race.Status,
	}
	if err := h.Service.CreateRace(raceToSave); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(raceToSave)
}

func (h *RaceHandler) GetRaceByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	race, err := h.Service.GetRaceByID(id)
	if err != nil || race == nil {
		http.Error(w, "Race not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(race)
}

func (h *RaceHandler) GetRaces(w http.ResponseWriter, r *http.Request) {
	races, err := h.Service.GetRaces()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(races)
}

func (h *RaceHandler) UpdateRace(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var race models.RaceRequest
	if err := json.NewDecoder(r.Body).Decode(&race); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if race.UserDNDID == 0 {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(race.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "user is not admin", http.StatusForbidden)
		return
	}

	existingRace, err := h.Service.GetRaceByID(id)
	if err != nil || existingRace == nil {
		http.Error(w, "Race not found", http.StatusNotFound)
		return
	}

	race.ID = id
	if race.Name != "" {
		existingRace.Name = race.Name
	}
	if race.Detail != "" {
		existingRace.Detail = race.Detail
	}
	if race.Status != "" {
		existingRace.Status = race.Status
	}
	if err := h.Service.UpdateRace(existingRace); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingRace)
}

func (h *RaceHandler) DeleteRace(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if r.Body == nil || r.Body == http.NoBody {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	var race models.RaceRequest
	if err := json.NewDecoder(r.Body).Decode(&race); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if race.UserDNDID == 0 {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(race.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "user is not admin", http.StatusForbidden)
		return
	}

	err = h.Service.DeleteRace(id)
	if err != nil {
		http.Error(w, "Race not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Race deleted successfully"})
}
