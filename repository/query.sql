-- name: GetPictureByID :one
SELECT *
FROM pictures
WHERE id = $1
LIMIT 1;

-- name: GetPicturesByFilter :many
SELECT *
FROM pictures
WHERE path LIKE '%' + $1 + '%'
LIMIT $2;

-- name: GetPicturesPaged :many
SELECT *
FROM pictures
ORDER BY updated DESC
OFFSET $1 LIMIT $2;

-- name: InsertPicture :exec
INSERT INTO pictures (path, ext, views, likes, rating, deviation, wins, losses, created, updated)
VALUES ($1, $2, 0, 0, 0, 0, 0, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT (path) DO UPDATE
    SET updated = CURRENT_TIMESTAMP;

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
ON CONFLICT (path) DO UPDATE
SET updated = CURRENT_TIMESTAMP
RETURNING *;

-- name: DeletePath :exec
DELETE
FROM paths
WHERE id = $1;
