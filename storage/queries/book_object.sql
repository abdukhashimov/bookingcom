-- name: GetBookObject :one
SELECT book.id,
    cat.name as "category",
    book.title,
    book.location,
    book.about,
    st.name as "status",
    book.opens_at,
    book.long,
    book.lat,
    book.closes_at,
    book.created_at,
    book.updated_at
FROM book_object as book
    LEFT JOIN category as cat ON book.category = cat.slug
    LEFT JOIN status as st ON book.status = st.id
WHERE book.id = $1
LIMIT 1;

-- name: GetAllBookObject :many
SELECT book.id,
    cat.name as "category",
    book.title,
    book.location,
    book.about,
    st.name as "status",
    book.opens_at,
    book.long,
    book.lat,
    book.closes_at,
    book.created_at,
    book.updated_at
FROM book_object as book
    LEFT JOIN category as cat ON book.category = cat.slug
    LEFT JOIN status as st ON book.status = st.id
WHERE status = $1
ORDER BY created_at desc OFFSET $2
LIMIT $3;

-- name: CreateBookObject :one
INSERT INTO book_object (
        id,
        category,
        title,
        location,
        about,
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
        $12
    )
RETURNING *;

-- name: UpdateBookObject :exec
UPDATE book_object
SET category = COALESCE(NULLIF(@category, ''), category),
    title = COALESCE(NULLIF(@title, ''), title),
    location = COALESCE(NULLIF(@location, ''), location),
    long = COALESCE(NULLIF(@long, 0), long),
    lat = COALESCE(NULLIF(@lat, 0), lat),
    about = COALESCE(NULLIF(@about, ''), about),
    status = COALESCE(NULLIF(@status, 0), status),
    opens_at = COALESCE(NULLIF(@opens_at, ''), opens_at),
    closes_at = COALESCE(NULLIF(@closes_at, ''), closes_at),
    updated_at = COALESCE(@updated_at, updated_at)
WHERE id = @id;

-- name: UpdateStatus :exec
UPDATE book_object
SET status = $1
WHERE id = $2;

-- name: DeleteBookObject :exec
DELETE FROM book_object
WHERE id = $1;
