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
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	DB *db.DB
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	userGroup := server.Group("/user")

	userGroup.PUT("/put", h.CreateUser)
	userGroup.GET("/get", h.GetUser)
	userGroup.DELETE("/delete", h.DeleteUser)
	userGroup.PATCH("/update", h.UpdateUser)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser models.User
	err := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %s", err))
		return
	}

	query := "INSERT INTO users (username, email) VALUES($1, $2) RETURNING id"
	err = h.DB.Conn.QueryRow(query, newUser.Username, newUser.Email).Scan(&newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating the user: %s", err))
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	var user models.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %s", err))
		return
	}

	query := "SELECT id, username, email FROM users WHERE id = $1"
	err = h.DB.Conn.QueryRow(query, user.ID).Scan(&user.ID, &user.Email, &user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error getting the user: %s", err))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var updatedUser models.User
	err := json.NewDecoder(c.Request.Body).Decode(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %s", err))
		return
	}
	if updatedUser.Username == "" || updatedUser.Email == "" {
		c.JSON(http.StatusBadRequest, "Both username and email are required")
		return
	}

	query := "UPDATE users SET username = $1, email = $2 WHERE id = $3"
	_, err = h.DB.Conn.Exec(query, updatedUser.Username, updatedUser.Email, updatedUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error updating the user: %s", err))
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var user models.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %s", err))
		return
	}
	log.Printf("deleting %d", user.ID)
	query := "DELETE FROM users WHERE id = $1"
	_, err = h.DB.Conn.Exec(query, user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("No user with id %d : %s", user.ID, err))
			return
		}
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error deleting the user: %s", err))
		return
	}

	c.JSON(http.StatusOK, user)
}
