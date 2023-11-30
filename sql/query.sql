-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (name, enabled, created_date, last_modified_date) VALUES ($1, $2, now(), now()) RETURNING *;

-- name: UpdateUser :one
UPDATE users set name = $2, enabled = $3, last_modified_date = now(), version = version + 1 WHERE id = $1 AND version = $4 RETURNING *;
