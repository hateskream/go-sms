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

-- name: GetAllSpaces :many
SELECT
    s.id,
    s.name,
    space_statuses.name as status,
    CAST(string_agg(f.name, ',') as VARCHAR) AS features
FROM
    spaces AS s
JOIN
    space_features sf ON s.id = sf.space_id
JOIN
    features f ON sf.feature_id = f.id    
JOIN
    (SELECT DISTINCT ON (s.id) s.id, space_statuses.name
     FROM spaces AS s
     JOIN space_statuses ON s.status_id = space_statuses.id
     ORDER BY s.id) AS space_statuses ON s.id = space_statuses.id
GROUP BY s.id, space_statuses.name;


-- name: GetSpacesByFeatureList :many
SELECT s.id, s.name
FROM 
  spaces s
JOIN 
  space_features sf ON s.id = sf.space_id
WHERE sf.feature_id = ANY(sqlc.arg(feature_list)::int[]) AND s.status_id = 1
GROUP BY s.id
HAVING  COUNT(DISTINCT sf.feature_id) = sqlc.arg(feature_count)::int; 