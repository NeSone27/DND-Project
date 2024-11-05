package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist-service/models"
	"todolist-service/services"

	"github.com/gorilla/mux"
)

type ClassHandler struct {
	Service     *services.ClassService
	ServiceUser *services.UserService
}

func NewClassHandler(classService *services.ClassService, userService *services.UserService) *ClassHandler {
	return &ClassHandler{
		Service:     classService,
		ServiceUser: userService,
	}
}

func (h *ClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var class models.ClassRequest
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if class.UserDNDID == 0 {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(class.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "user is not admin", http.StatusForbidden)
		return
	}

	classToSave := &models.Class{
		Name:   class.Name,
		Detail: class.Detail,
		Status: class.Status,
	}
	if err := h.Service.CreateClass(classToSave); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

func (h *ClassHandler) GetClassByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	class, err := h.Service.GetClassByID(id)
	if err != nil || class == nil {
		http.Error(w, "Class not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(class)
}

func (h *ClassHandler) GetClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := h.Service.GetClasses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(classes)
}

func (h *ClassHandler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var class models.ClassRequest
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if class.UserDNDID == 0 {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(class.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "user is not admin", http.StatusForbidden)
		return
	}

	existingClass, err := h.Service.GetClassByID(id)
	if err != nil || existingClass == nil {
		http.Error(w, "Class not found", http.StatusNotFound)
		return
	}

	class.ID = id
	if class.Name != "" {
		existingClass.Name = class.Name
	}
	if class.Detail != "" {
		existingClass.Detail = class.Detail
	}
	if class.Status != "" {
		existingClass.Status = class.Status
	}
	if err := h.Service.UpdateClass(existingClass); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingClass)
}

func (h *ClassHandler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if r.Body == nil || r.Body == http.NoBody {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	var class models.ClassRequest
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if class.UserDNDID == 0 {
		http.Error(w, "user_dnd_id is required", http.StatusBadRequest)
		return
	}
	user, err := h.ServiceUser.GetUserByID(class.UserDNDID)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "user is not admin", http.StatusForbidden)
		return
	}

	err = h.Service.DeleteClass(id)
	if err != nil {
		http.Error(w, "Class not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Class deleted successfully"})
}
