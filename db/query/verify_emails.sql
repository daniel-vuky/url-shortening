-- name: CreateVerifyEmailToken :one
INSERT INTO verify_emails (email, token, expires_at) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateVerifyEmailToken :one
UPDATE verify_emails SET is_used = true WHERE email = $1 AND token = $2 AND expires_at > NOW() RETURNING *;