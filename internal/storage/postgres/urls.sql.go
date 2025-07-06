package postgres

import (
	"context"
	"time"

	"github.com/daniel-vuky/url-shortening/internal/models"
)

const createURL = `-- name: CreateURL :one
INSERT INTO urls (original_url, short_code, expires_at) VALUES ($1, $2, $3) RETURNING id, original_url, short_code, created_at, expires_at, total_click, last_click_at, unique_vistor, user_id
`

type CreateURLParams struct {
	OriginalUrl string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func (q *Queries) CreateURL(ctx context.Context, arg CreateURLParams) (models.URL, error) {
	row := q.db.QueryRow(ctx, createURL, arg.OriginalUrl, arg.ShortCode, arg.ExpiresAt)
	var i models.URL
	err := row.Scan(
		&i.ID,
		&i.OriginalURL,
		&i.ShortCode,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.TotalClick,
		&i.LastClickAt,
		&i.UniqueVistor,
		&i.UserID,
	)
	return i, err
}

const deleteURL = `-- name: DeleteURL :one
DELETE FROM urls WHERE id = $1 RETURNING id, original_url, short_code, created_at, expires_at, total_click, last_click_at, unique_vistor, user_id
`

func (q *Queries) DeleteURL(ctx context.Context, id int32) (models.URL, error) {
	row := q.db.QueryRow(ctx, deleteURL, id)
	var i models.URL
	err := row.Scan(
		&i.ID,
		&i.OriginalURL,
		&i.ShortCode,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.TotalClick,
		&i.LastClickAt,
		&i.UniqueVistor,
		&i.UserID,
	)
	return i, err
}

const getListURL = `-- name: GetListURL :many
SELECT id, original_url, short_code, created_at, expires_at, total_click, last_click_at, unique_vistor, user_id FROM urls WHERE user_id = $1 LIMIT $2 OFFSET $3
`

type GetListURLParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListURL(ctx context.Context, arg GetListURLParams) ([]models.URL, error) {
	rows, err := q.db.Query(ctx, getListURL, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.URL{}
	for rows.Next() {
		var i models.URL
		if err := rows.Scan(
			&i.ID,
			&i.OriginalURL,
			&i.ShortCode,
			&i.CreatedAt,
			&i.ExpiresAt,
			&i.TotalClick,
			&i.LastClickAt,
			&i.UniqueVistor,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getURL = `-- name: GetURL :one
SELECT id, original_url, short_code, created_at, expires_at, total_click, last_click_at, unique_vistor, user_id FROM urls WHERE id = $1
`

func (q *Queries) GetURL(ctx context.Context, id int32) (models.URL, error) {
	row := q.db.QueryRow(ctx, getURL, id)
	var i models.URL
	err := row.Scan(
		&i.ID,
		&i.OriginalURL,
		&i.ShortCode,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.TotalClick,
		&i.LastClickAt,
		&i.UniqueVistor,
		&i.UserID,
	)
	return i, err
}

const updateURL = `-- name: UpdateURL :one
UPDATE urls
SET
    original_url = COALESCE($2, original_url),
    short_code = COALESCE($3, short_code),
    expires_at = COALESCE($4, expires_at),
    total_click = COALESCE($5, total_click),
    last_click_at = COALESCE($6, last_click_at),
    unique_vistor = COALESCE($7, unique_vistor)
WHERE id = $1
RETURNING id, original_url, short_code, created_at, expires_at, total_click, last_click_at, unique_vistor, user_id
`

type UpdateURLParams struct {
	ID           int32     `json:"id"`
	OriginalUrl  string    `json:"original_url"`
	ShortCode    string    `json:"short_code"`
	ExpiresAt    time.Time `json:"expires_at"`
	TotalClick   int64     `json:"total_click"`
	LastClickAt  time.Time `json:"last_click_at"`
	UniqueVistor int64     `json:"unique_vistor"`
}

func (q *Queries) UpdateURL(ctx context.Context, arg UpdateURLParams) (models.URL, error) {
	row := q.db.QueryRow(ctx, updateURL,
		arg.ID,
		arg.OriginalUrl,
		arg.ShortCode,
		arg.ExpiresAt,
		arg.TotalClick,
		arg.LastClickAt,
		arg.UniqueVistor,
	)
	var i models.URL
	err := row.Scan(
		&i.ID,
		&i.OriginalURL,
		&i.ShortCode,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.TotalClick,
		&i.LastClickAt,
		&i.UniqueVistor,
		&i.UserID,
	)
	return i, err
}
