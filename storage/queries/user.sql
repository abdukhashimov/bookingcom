-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByPhoneNumber :one
SELECT * from users WHERE phone_number = $1 LIMIT 1;

-- name: GetUsers :many
SELECT *
FROM users
ORDER BY created_at desc OFFSET $1
LIMIT $2;

-- name: CreateUser :one
INSERT INTO users (
        id,
        first_name,
        last_name,
        phone_number,
        is_verified,
        long,
        lat,
        user_type,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET first_name = $1,
    last_name = $2,
    phone_number = $3,
    is_verified = $4,
    long = $5,
    lat = $6,
    user_type = $7,
    updated_at = $8
WHERE id = $9;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;