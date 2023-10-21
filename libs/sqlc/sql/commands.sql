-- name: GetAnswer :one
SELECT answer FROM commands
         WHERE channel_id = $1
         AND command = $2;
