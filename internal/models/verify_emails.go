package models

import (
	"time"
)

type VerifyEmail struct {
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	IsUsed    bool      `json:"is_used"`
	ExpiresAt time.Time `json:"expires_at"`
}
