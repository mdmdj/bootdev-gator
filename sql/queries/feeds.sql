-- name: CreateFeed :one
INSERT INTO feeds (url, user_id, name)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds LEFT JOIN users ON users.id = feeds.user_id ORDER BY feeds.created_at DESC;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;