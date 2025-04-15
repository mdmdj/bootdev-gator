-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2
)
RETURNING *;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users ORDER BY name ASC;

-- name: ResetUsers :exec
TRUNCATE TABLE users CASCADE;