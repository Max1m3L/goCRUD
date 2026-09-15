package handlers

import (
	"cosmic-crud/models"
	"cosmic-crud/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type SpaceshipHandler struct {
	service *services.SpaceshipService
}

func NewSpaceshipHandler(service *services.SpaceshipService) *SpaceshipHandler {
	return &SpaceshipHandler{service: service}
}

// POST /spaceships
func (h *SpaceshipHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ship models.Spaceship
	err := json.NewDecoder(r.Body).Decode(&ship)
	if err != nil {
		http.Error(w, "Неверный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Create(&ship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ship)
}

// GET /spaceships
func (h *SpaceshipHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ships, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ships)
}

// GET /spaceships/{id}
func (h *SpaceshipHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	ship, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ship)
}

// PUT /spaceships/{id}
func (h *SpaceshipHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var ship models.Spaceship
	err = json.NewDecoder(r.Body).Decode(&ship)
	if err != nil {
		http.Error(w, "Неверный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	ship.ID = id

	err = h.service.Update(&ship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ship)
}

// DELETE /spaceships/{id}
func (h *SpaceshipHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// POST /spaceships/{id}/refuel
func (h *SpaceshipHandler) Refuel(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Amount int `json:"amount"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Неверный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	ship, err := h.service.Refuel(id, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ship)
}

// Вспомогательная функция для извлечения ID из URL
func extractID(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("неверный формат URL")
	}
	return strconv.Atoi(parts[2])
}
