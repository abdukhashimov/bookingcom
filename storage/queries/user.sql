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
SET first_name = COALESCE($1, first_name),
    last_name = COALESCE($2, last_name),
    phone_number = COALESCE($3, phone_number),
    is_verified = COALESCE($4, is_verified),
    long = COALESCE($5, long),
    lat = COALESCE($6, lat),
    user_type = COALESCE($7, user_type),
    updated_at = COALESCE($8, updated_at)
WHERE id = $9;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;