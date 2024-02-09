-- name: AddSpace :one
INSERT INTO spaces (name, physical_id, group_id, status_id )
VALUES ($1,
CASE WHEN $2 = '' THEN NULL ELSE $2,
CASE WHEN $3 = '' THEN NULL ELSE $3,
CASE WHEN $4 = '' THEN NULL ELSE $4);

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

-- name: UpdateSpaceStatus :exec
UPDATE spaces
  set status_id = $3
WHERE spaces.id = $1 AND spaces.status_id = $2;

-- name: GetCardByNumber :one
SELECT cars.id
FROM cars
WHERE cars.number = sqlc.arg(number)::VARCHAR;

-- name: AddCardNumber :one
INSERT INTO cars (number)
VALUES (sqlc.arg(number)::VARCHAR) RETURNING id;

-- name: AddReservation :one
INSERT INTO space_reservations (time_from, time_to, car_id, reservation_fee, space_id, status_id)
VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;

-- name: UpdateReservationState :exec
UPDATE space_reservations
  set status_id = $2
WHERE space_reservations.id = $1;

-- name: GetSpaceStatuses :many
Select *
From space_statuses;

