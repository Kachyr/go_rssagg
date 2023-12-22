package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"rssagg/internal/database"
	"time"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, dbErr := apiCfg.DB.CreateFeed(
		r.Context(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      params.Name,
			Url:       params.URL,
			UserID:    user.ID,
		})

	if dbErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", dbErr))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, dbErr := apiCfg.DB.GetFeeds(
		r.Context(),
	)
	if dbErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", dbErr))
		return
	}
	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
