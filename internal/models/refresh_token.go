package models

import "time"

type RefreshToken struct {
	RefreshTokenID int64     `json:"refresh_token_id"`
	UserID         int64     `json:"user_id"`
	RefreshToken   string    `json:"refresh_token"`
	UserAgent      string    `json:"user_agent"`
	ClientIp       string    `json:"client_ip"`
	IsBlocked      bool      `json:"is_blocked"`
	ExpiredAt      time.Time `json:"expired_at"`
	CreatedAt      time.Time `json:"created_at"`
}
