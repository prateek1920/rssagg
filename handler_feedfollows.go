package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "github.com/prateek1920/rssagg/internal/database"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error occurred in reading params: %v", err))
		return
	}

	feed_follows, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error occurred in creating feed: %v", err))
		return
	}

	responseWithJSON(w, 200, DatabaseFeedFollowToFeedFollow(feed_follows))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feed_follows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error occurred in fetching feed_follows: %v", err))
		return
	}

	responseWithJSON(w, 200, DatabaseFeedFollowsToFeedFollows(feed_follows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowStr := chi.URLParam(r, "feedFollowsID")
	feedFollowID, err := uuid.Parse(feedFollowStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error occurred in parsing feed_follows id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error occurred in deleting feed_follows: %v", err))
		return
	}

	responseWithJSON(w, 200, struct{}{})
}
