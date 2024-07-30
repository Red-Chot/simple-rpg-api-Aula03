package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	 "github.com/Red-Chot/simple-rpg-api-Aula03/internal/entity"
	"github.com/Red-Chot/simple-rpg-api-Aula03/internal/service"
)

type EnemyHandler struct {
	EnemyService *service.EnemyService
}

 func NewEnemyHandler(enemyService *service.EnemyService) *EnemyHandler {
	return &EnemyHandler{EnemyService: enemyService}
}

 func (eh *EnemyHandler) AddEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		Nickname string `json:"nickname"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: "internal server error"})
		return
	}

	enemy, err := eh.EnemyService.AddEnemy(request.Nickname)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemy)
}

func (eh *EnemyHandler) LoadEnemies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	enemies, err := eh.EnemyService.LoadEnemies()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemies)
}

func (eh *EnemyHandler) DeleteEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Path[len("/enemy/"):]

	if err := eh.EnemyService.DeleteEnemy(id); err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}

func (eh *EnemyHandler) LoadEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Path[len("/enemy/"):]

	enemy, err := eh.EnemyService.LoadEnemy(id)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemy)
}

func (eh *EnemyHandler) SaveEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Path[len("/enemy/"):]

	var request struct {
		Nickname string `json:"nickname"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: "internal server error"})
		return
	}

	enemy, err := eh.EnemyService.SaveEnemy(id, request.Nickname)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemy)
}































