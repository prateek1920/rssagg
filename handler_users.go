package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "/Users/prateekgupta/rssagg/internal/database/db.go"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error occurred in reading params: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.User{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error occurred in creating user: %v", err))
		return
	}

	responseWithJSON(w, 200, user)
}
