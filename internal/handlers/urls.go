package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/daniel-vuky/url-shortening/internal/models"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) CreateURL(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	expiresAtParam := r.FormValue("expires_at")
	var expiresAt *time.Time
	if expiresAtParam != "" {
		expiresAtTime, err := time.Parse(time.RFC3339, expiresAtParam)
		if err != nil {
			http.Error(w, "Invalid expires_at format", http.StatusBadRequest)
			return
		}
		expiresAt = &expiresAtTime
	}
	userIDParam := r.FormValue("user_id")
	var userID *string
	if userIDParam != "" {
		userID = &userIDParam
	}
	requestParam := &models.CreateURLRequest{
		OriginalURL: r.FormValue("url"),
		ShortCode:   r.FormValue("short_code"),
		ExpiresAt:   expiresAt,
		UserID:      userID,
	}

	if err := validator.New().Struct(requestParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.service.CreateURL(requestParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
