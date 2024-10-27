-- name: GetUser :one
SELECT * FROM "user" WHERE id = $1 LIMIT 1;

-- name: PostUser :one
INSERT INTO "user" (id, name, icon_url) VALUES ($1, $2, $3) RETURNING *;

-- name: GetCassette :one
SELECT * FROM cassette WHERE id = $1 LIMIT 1;

-- name: GetCassettesByUser :many
SELECT * FROM cassette WHERE user_id = $1;

-- name: PostCassette :one
INSERT INTO cassette (user_id, name) VALUES ($1, $2) RETURNING *;

-- name: PostSong :one
INSERT INTO songs (id, cassette_id, user_id, song_number, ,song_time, name, url, upload_user) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;
