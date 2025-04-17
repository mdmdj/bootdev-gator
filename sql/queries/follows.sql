-- name: CreateFeedFollow :one
WITH new_feed_follow AS (
     INSERT INTO feed_follows (user_id, feed_id)
     VALUES ($1, $2) RETURNING *
)
SELECT
    new_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM new_feed_follow
INNER JOIN users ON users.id = new_feed_follow.user_id
INNER JOIN feeds ON feeds.id = new_feed_follow.feed_id;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows ORDER BY created_at DESC;

-- name: GetFeedFollowsByUser :many
SELECT * FROM feed_follows WHERE user_id = $1 ORDER BY created_at DESC;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = $1;
