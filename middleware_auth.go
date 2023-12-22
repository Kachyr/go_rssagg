package main

import (
	"fmt"
	"net/http"
	"rssagg/auth"
	"rssagg/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, dbErr := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if dbErr != nil {
			respondWithError(w, 401, fmt.Sprintf("Couldn't get user: %v", dbErr))
			return
		}

		handler(w, r, user)

	}
}
