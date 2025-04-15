-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, url, user_id, name)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds LEFT JOIN users ON users.id = feeds.user_id ORDER BY feeds.created_at DESC;