-- name: AddSpace :one
INSERT INTO spaces (name)
VALUES ($1) RETURNING *;

-- name: AddSpaceFeature :exec
INSERT INTO space_features (space_id,Feature_id)
VALUES ($1,$2);

-- name: GetSpaces :many
SELECT *
FROM spaces
ORDER BY id;

-- name: DeleteSpace :exec
DELETE FROM spaces
WHERE id = $1;

-- name: GetFeatures :many
SELECT *
FROM features
ORDER BY id;

-- name: AddFeature :one
INSERT INTO features (name)
VALUES ($1) RETURNING id;

-- name: UpdateFeature :exec
UPDATE features
  set name = $2  
WHERE id = $1;

-- name: DeleteFeature :one
DELETE FROM features
WHERE id = $1 RETURNING id;