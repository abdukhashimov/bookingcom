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
        category_id,
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
SET
    category_id = COALESCE(@category_id, category_id),
    title = COALESCE(@title, title),
    location = COALESCE(@location, location),
    long = COALESCE(@long, long),
    lat = COALESCE(@lat, lat),
    about = COALESCE(@about, about),
    discount = COALESCE(@discount, discount),
    discount_expires = COALESCE(@discount_expires, discount_expires),
    status = COALESCE(@status, status),
    opens_at = COALESCE(@opens_at, opens_at),
    closes_at = COALESCE(@closes_at, closes_at),
    updated_at = COALESCE(@updated_at, updated_at)
WHERE id = @id;

-- name: DeleteBookObject :exec
DELETE FROM book_object
WHERE id = $1;