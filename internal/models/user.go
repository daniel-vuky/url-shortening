package models

import (
	"time"
)

type User struct {
	UserID         int32     `json:"user_id"`
	Email          string    `json:"email"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	EmailVerified  string    `json:"email_verified"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
}
