-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, hashed_password)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2)
RETURNING *;

-- name: GetUserFromUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: ChangePasswordAndUsername :one
UPDATE users SET username = $2,
hashed_password = $3
WHERE id = $1
RETURNING *;