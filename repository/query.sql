-- name: GetPictureByID :one
SELECT *
FROM pictures
WHERE id = $1
LIMIT 1;

-- name: GetPictureByFilter :one
SELECT *
FROM pictures
WHERE path LIKE '%' + $1 + '%'
LIMIT 1;

-- name: GetPicturesByFilter :many
SELECT *
FROM pictures
WHERE path LIKE '%' + $1 + '%';

-- name: GetPathByID :one
SELECT *
FROM paths
WHERE id = $1
LIMIT 1;

-- name: GetPaths :many
SELECT *
FROM paths
ORDER BY path;

-- name: CreatePath :one
INSERT INTO paths (path, created, updated)
VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: DeletePath :exec
DELETE
FROM paths
WHERE id = $1;
