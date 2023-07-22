package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "github.com/prateek1920/rssagg/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error occurred in reading params: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      user.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error occurred in creating feed: %v", err))
		return
	}

	responseWithJSON(w, 200, DatabaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error occurred in fetching feeds: %v", err))
		return
	}

	responseWithJSON(w, 200, DatabaseFeedsToFeeds(feeds))
}
