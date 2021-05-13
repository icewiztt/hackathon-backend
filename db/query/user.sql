-- name: CreateUser :one
INSERT INTO users (
    username,
    fullname,
    password_encoded,
    usertype
) VALUES (
    $1 , $2 , $3 , $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users 
ORDER BY username 
LIMIT $1
OFFSET $2;

-- name: UpdateUserPassword :one
UPDATE users
SET password_encoded = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserFullname :one
UPDATE users
SET fullname = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;