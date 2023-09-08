-- name: CreateChannel :one
INSERT INTO channels (
    twitch_name, twitch_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetChannel :one
SELECT * FROM channels
         WHERE id = $1;

-- name: GetChannels :many
SELECT * FROM channels;

-- name: DeleteChannel :exec
DELETE FROM channels
       WHERE id = $1;

-- name: UpdateChannel :one
UPDATE channels
   SET twitch_name = $2, twitch_id = $3
 WHERE id = $1
 RETURNING *;