package postgres

import (
	"context"
	"time"

	"github.com/daniel-vuky/url-shortening/internal/models"
)

const createVerifyEmailToken = `-- name: CreateVerifyEmailToken :one
INSERT INTO verify_emails (email, token, expires_at) VALUES ($1, $2, $3) RETURNING email, token, is_used, expires_at
`

type CreateVerifyEmailTokenParams struct {
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (q *Queries) CreateVerifyEmailToken(ctx context.Context, arg CreateVerifyEmailTokenParams) (models.VerifyEmail, error) {
	row := q.db.QueryRow(ctx, createVerifyEmailToken, arg.Email, arg.Token, arg.ExpiresAt)
	var i models.VerifyEmail
	err := row.Scan(
		&i.Email,
		&i.Token,
		&i.IsUsed,
		&i.ExpiresAt,
	)
	return i, err
}

const updateVerifyEmailToken = `-- name: UpdateVerifyEmailToken :one
UPDATE verify_emails SET is_used = true WHERE email = $1 AND token = $2 AND expires_at > NOW() RETURNING email, token, is_used, expires_at
`

type UpdateVerifyEmailTokenParams struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (q *Queries) UpdateVerifyEmailToken(ctx context.Context, arg UpdateVerifyEmailTokenParams) (models.VerifyEmail, error) {
	row := q.db.QueryRow(ctx, updateVerifyEmailToken, arg.Email, arg.Token)
	var i models.VerifyEmail
	err := row.Scan(
		&i.Email,
		&i.Token,
		&i.IsUsed,
		&i.ExpiresAt,
	)
	return i, err
}
