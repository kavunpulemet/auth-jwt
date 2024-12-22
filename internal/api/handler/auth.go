package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"auth/internal/models"
	"auth/internal/service"
	"auth/internal/utils"
)

func SignIn(ctx utils.MyContext, service service.AuthorizationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["userID"]
		if userID == "" {
			utils.NewErrorResponse(ctx, w, "userID cannot be empty", http.StatusBadRequest)
			return
		}

		_, err := uuid.Parse(userID)
		if err != nil {
			utils.NewErrorResponse(ctx, w, "invalid userID format", http.StatusBadRequest)
			return
		}

		userIP := utils.GetClientIP(r)
		if userIP == "" {
			utils.NewErrorResponse(ctx, w, "unable to determine user IP", http.StatusBadRequest)
			return
		}

		tokenPair, err := service.SignIn(ctx, userID, userIP)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = utils.WriteResponse(w, http.StatusOK, tokenPair); err != nil {
			utils.NewErrorResponse(ctx, w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func Refresh(ctx utils.MyContext, service service.AuthorizationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tokenPair models.TokenPair
		if err := json.NewDecoder(r.Body).Decode(&tokenPair); err != nil {
			utils.NewErrorResponse(ctx, w, "invalid JSON payload", http.StatusBadRequest)
			return
		}

		if tokenPair.AccessToken == "" || tokenPair.RefreshToken == "" {
			utils.NewErrorResponse(ctx, w, "both access and refresh tokens are required", http.StatusBadRequest)
			return
		}

		userIP := utils.GetClientIP(r)
		if userIP == "" {
			utils.NewErrorResponse(ctx, w, "unable to determine user IP", http.StatusBadRequest)
			return
		}

		newTokenPair, err := service.RefreshTokens(ctx, tokenPair.AccessToken, tokenPair.RefreshToken, userIP)
		if err != nil {
			if err.Error() == "hash and token do not match" || err.Error() == "invalid access token" {
				utils.NewErrorResponse(ctx, w, err.Error(), http.StatusUnauthorized)
			} else {
				utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if err = utils.WriteResponse(w, http.StatusOK, newTokenPair); err != nil {
			utils.NewErrorResponse(ctx, w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}
