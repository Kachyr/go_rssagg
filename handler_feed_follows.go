package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"rssagg/internal/database"
	"time"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, dbErr := apiCfg.DB.CreateFeedFollows(
		r.Context(),
		database.CreateFeedFollowsParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			FeedID:    params.FeedId,
		})

	if dbErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", dbErr))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, dbErr := apiCfg.DB.GetFeedFollows(
		r.Context(),
		user.ID,
	)
	if dbErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get User feeds: %v", dbErr))
		return
	}
	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feeds))
}

func (apiCfg *apiConfig) handleDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, prsErr := uuid.Parse(feedFollowIdStr)

	if prsErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed_id: %v", prsErr))
		return
	}

	dbErr := apiCfg.DB.DeleteFeedFollows(
		r.Context(),

		database.DeleteFeedFollowsParams{ID: feedFollowId, UserID: user.ID},
	)
	if dbErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", dbErr))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
