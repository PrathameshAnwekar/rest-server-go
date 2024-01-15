package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PrathameshAnwekar/rest-server-go/db"
	"github.com/PrathameshAnwekar/rest-server-go/models"
)

type UserHandler struct {
	DB *db.DB
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %s", err), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO users (username, email) VALUES($1, $2) RETURNING id"
	err = h.DB.Conn.QueryRow(query, newUser.Username, newUser.Email).Scan(&newUser.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating the user: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error returning the user: %s", err), http.StatusInternalServerError)
		return
	}
}

// func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	
// }

// func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

// }

// func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

// }
