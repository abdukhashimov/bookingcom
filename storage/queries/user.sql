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
SET first_name = COALESCE(@first_name, first_name),
    last_name = COALESCE(@last_name, last_name),
    phone_number = COALESCE(NULLIF(@phone_number, ''), phone_number),
    is_verified = COALESCE(@is_verified, is_verified),
    long = COALESCE(@long, long),
    lat = COALESCE(@lat, lat),
    user_type = COALESCE(@user_type, user_type),
    updated_at = COALESCE(@updated_at, updated_at)
WHERE id = @id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;