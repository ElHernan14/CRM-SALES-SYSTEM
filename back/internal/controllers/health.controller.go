package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type HealthHandler struct {
	DB *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{DB: db}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	err := h.DB.Ping()

	response := map[string]interface{}{
		"status": "ok",
		"db":     "up",
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response["status"] = "error"
		response["db"] = "down"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
