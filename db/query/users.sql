-- name: GetUser :one
SELECT * FROM users where email = $1;

-- name: CreateUser :one
INSERT INTO users (email, firstname, lastname, hashed_password, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    firstname = COALESCE(sqlc.narg(firstname), firstname),
    lastname = COALESCE(sqlc.narg(lastname), lastname),
    email_verified = COALESCE(sqlc.narg(email_verified), email_verified),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password)
WHERE email = $1
RETURNING *;