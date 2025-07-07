package postgres

import (
	"context"
	"time"

	"github.com/daniel-vuky/url-shortening/internal/models"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, firstname, lastname, hashed_password, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING user_id, email, firstname, lastname, email_verified, hashed_password, created_at
`

type CreateUserParams struct {
	Email          string    `json:"email"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (models.User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.HashedPassword,
		arg.CreatedAt,
	)
	var i models.User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.EmailVerified,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT user_id, email, firstname, lastname, email_verified, hashed_password, created_at FROM users where email = $1
`

func (q *Queries) GetUser(ctx context.Context, email string) (models.User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i models.User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.EmailVerified,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    firstname = COALESCE($2, firstname),
    lastname = COALESCE($3, lastname),
    email_verified = COALESCE($4, email_verified),
    hashed_password = COALESCE($5, hashed_password)
WHERE email = $1
RETURNING user_id, email, firstname, lastname, email_verified, hashed_password, created_at
`

type UpdateUserParams struct {
	Email          string `json:"email"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	EmailVerified  bool   `json:"email_verified"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (models.User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.EmailVerified,
		arg.HashedPassword,
	)
	var i models.User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.EmailVerified,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}
