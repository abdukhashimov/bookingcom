-- name: GetBookObject :one
SELECT *
FROM book_object
WHERE id = $1
LIMIT 1;

-- name: GetAllBookObject :many
SELECT *
FROM book_object
ORDER BY created_at desc OFFSET $1
LIMIT $2;

-- name: CreateBookObject :one
INSERT INTO book_object (
        id,
        category,
        title,
        location,
        about,
        discount,
        discount_expires,
        status,
        opens_at,
        long,
        lat,
        closes_at,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14
    )
RETURNING *;

-- name: UpdateBookObject :exec
UPDATE book_object
SET category = COALESCE(NULLIF(@category, ''), category),
    title = COALESCE(NULLIF(@title, ''), title),
    location = COALESCE(NULLIF(@location, ''), location),
    long = COALESCE(@long, long),
    lat = COALESCE(@lat, lat),
    about = COALESCE(NULLIF(@about, ''), about),
    discount = COALESCE(@discount, discount),
    discount_expires = COALESCE(NULLIF(@discount_expires, ''), discount_expires),
    status = COALESCE(NULLIF(@status, ''), status),
    opens_at = COALESCE(NULLIF(@opens_at, ''), opens_at),
    closes_at = COALESCE(NULLIF(@closes_at, ''), closes_at),
    updated_at = COALESCE(@updated_at, updated_at)
WHERE id = @id;

-- name: DeleteBookObject :exec
DELETE FROM book_object
WHERE id = $1;