-- name: CreateURL :one
INSERT INTO urls (original_url, short_code, expires_at) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateURL :one
UPDATE urls
SET
    original_url = COALESCE(sqlc.narg(original_url), original_url),
    short_code = COALESCE(sqlc.narg(short_code), short_code),
    expires_at = COALESCE(sqlc.narg(expires_at), expires_at),
    total_click = COALESCE(sqlc.narg(total_click), total_click),
    last_click_at = COALESCE(sqlc.narg(last_click_at), last_click_at),
    unique_vistor = COALESCE(sqlc.narg(unique_vistor), unique_vistor)
WHERE id = $1
RETURNING *;

-- name: DeleteURL :one
DELETE FROM urls WHERE id = $1 RETURNING *;

-- name: GetURL :one
SELECT * FROM urls WHERE id = $1;

-- name: GetListURL :many
SELECT * FROM urls WHERE user_id = $1 LIMIT $2 OFFSET $3;