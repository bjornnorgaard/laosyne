-- name: GetPictureByID :one
SELECT *
FROM pictures
WHERE id = $1 LIMIT 1;

-- name: GetPictureByFilter :one
SELECT *
FROM pictures
WHERE path LIKE '%' + $1 + '%' LIMIT 1;

-- name: GetPicturesByFilter :many
SELECT *
FROM pictures
WHERE path LIKE '%' + $1 + '%';

-- name: GetMediaPathByID :one
SELECT *
FROM media_paths
WHERE id = $1 LIMIT 1;

-- name: GetMediaPaths :many
SELECT *
FROM media_paths
ORDER BY path;

-- name: CreateMediaPath :one
INSERT INTO media_paths (path)
VALUES ($1) RETURNING *;
