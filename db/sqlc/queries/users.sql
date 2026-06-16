-- name: CreateUser :one
-- result: last_insert_id
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING id;

-- name: GetUserByID :one
-- one:
SELECT id, name, dob, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $1,
    dob = $2,
    updated_at = NOW()
WHERE id = $3;
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many 
-- many:
SELECT id, name, dob, created_at, updated_at
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;
