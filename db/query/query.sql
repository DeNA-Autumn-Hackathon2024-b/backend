-- name: GetUser :one
SELECT * FROM "user" WHERE id = $1 LIMIT 1;

-- name: PostUser :one
INSERT INTO "user" (id, name, icon_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetCassette :one
SELECT * FROM cassette WHERE id = $1 LIMIT 1;

-- name: GetCassettesByUser :many
SELECT * FROM cassette WHERE user_id = $1;

-- name: PostCassette :one
INSERT INTO cassette (id, user_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;
