-- name: GetStatusByName :one
SELECT * FROM status WHERE name = $1;