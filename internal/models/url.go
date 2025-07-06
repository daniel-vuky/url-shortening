package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type URL struct {
	ID           int32              `json:"id"`
	OriginalURL  string             `json:"original_url"`
	ShortCode    string             `json:"short_code"`
	CreatedAt    time.Time          `json:"created_at"`
	ExpiresAt    pgtype.Timestamptz `json:"expires_at"`
	TotalClick   int32              `json:"total_click"`
	LastClickAt  pgtype.Timestamptz `json:"last_click_at"`
	UniqueVistor int32              `json:"unique_vistor"`
	UserID       pgtype.Int8        `json:"user_id"`
}

type CreateURLRequest struct {
	OriginalURL string     `json:"original_url"`
	CustomCode  string     `json:"custom_code"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	UserID      *string    `json:"user_id,omitempty"`
}

type CreateURLResponse struct {
	ShortURL    string     `json:"short_url"`
	OriginalURL string     `json:"original_url"`
	ShortCode   string     `json:"short_code"`
	CreatedAt   time.Time  `json:"created_at"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	UsedCount   int64      `json:"used_count"`
}

type URLStats struct {
	TotalClick    int64      `json:"total_click"`
	LastClickAt   *time.Time `json:"last_click_at,omitempty"`
	UniqueVistors int64      `json:"unique_vistors"`
	CreatedAt     time.Time  `json:"created_at"`
}
