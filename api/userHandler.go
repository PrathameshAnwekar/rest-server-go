package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

	setUserHandlerResponse(&newUser, w)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %s", err), http.StatusBadRequest)
		return
	}

	query := "SELECT id, username, email FROM users WHERE id = $1"
	err = h.DB.Conn.QueryRow(query, user.ID).Scan(&user.ID, &user.Email, &user.Username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting the user: %s", err), http.StatusInternalServerError)
		return
	}

	setUserHandlerResponse(&user, w)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %s", err), http.StatusBadRequest)
		return
	}
	if updatedUser.Username == "" || updatedUser.Email == "" {
		http.Error(w, "Both username and email are required", http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET username = $1, email = $2 WHERE id = $3"
	_, err = h.DB.Conn.Exec(query, updatedUser.Username, updatedUser.Email, updatedUser.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating the user: %s", err), http.StatusInternalServerError)
		return
	}

	setUserHandlerResponse(&updatedUser, w)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %s", err), http.StatusBadRequest)
		return
	}
	log.Printf("deleting %d", user.ID)
	query := "DELETE FROM users WHERE id = $1"
	_, err = h.DB.Conn.Exec(query, user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, fmt.Sprintf("No user with id %d : %s", user.ID, err), http.StatusInternalServerError)
			return
		}
		http.Error(w, fmt.Sprintf("Error deleting the user: %s", err), http.StatusInternalServerError)
		return
	}

	setUserHandlerResponse(&user, w)
}

func setUserHandlerResponse(user *models.User, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error returning the user: %s", err), http.StatusInternalServerError)
		return
	}
}
