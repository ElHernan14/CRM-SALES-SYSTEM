package users

import (
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := []map[string]string{
		{"id": "1", "name": "Hernan"},
		{"id": "2", "name": "Juan"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	users := []map[string]string{
		{"id": "1", "name": "Hernan"},
		{"id": "2", "name": "Juan"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
