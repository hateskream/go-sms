-- name: AddSpace :one
INSERT INTO spaces (name, physical_id, group_id, status_id)
VALUES ($1,$2,$3,$4)
RETURNING  id;


-- name: GetPricingPolicy :many
SELECT pg.name, tpp.*
FROM time_pricing_policy as tpp
JOIN pricing_groups as pg ON pg.id = tpp.group_id
WHERE pg.id = $1;

-- name: AddPricingGroup :one
INSERT INTO pricing_groups (name)
VALUES ($1) RETURNING id;

-- name: GeneratTimePricingPolicy :exec
INSERT INTO time_pricing_policy (rate, hour, day_of_week, group_id)
SELECT $2, hour, day_of_week, $1
FROM generate_series(1, 8) AS day_of_week
CROSS JOIN generate_series(1, 24) AS hour;

-- name: DeletePricingGroup :one
DELETE FROM pricing_groups
WHERE id = $1 RETURNING id;

-- name: CleanTimePricingPolicy :exec
DELETE FROM time_pricing_policy
WHERE group_id = $1;


-- name: UpdatePricingGroups :exec
UPDATE pricing_groups
  set name = $2
  WHERE id = $1;

-- name: UpdatePricingPolicy :exec
UPDATE time_pricing_policy
  set rate = $2
  WHERE ID = $1;


-- name: DeleteSpace :one
DELETE FROM spaces
WHERE id = $1 RETURNING id;

-- name: UpdateSpace :exec
UPDATE spaces
  set name = $2,
  physical_id = $3,
  group_id = $4,
  status_id = $5
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

-- name: GetSpaces :many
SELECT
    s.id,
    s.name,
    space_statuses.name AS status,        
    COALESCE(CAST(string_agg(CASE WHEN sf.is_required THEN f.name END, ',') AS VARCHAR), '') AS required_features,
    COALESCE(CAST(string_agg(f.name, ',') AS VARCHAR), '') AS features
FROM
    spaces AS s
JOIN
    space_features sf ON s.id = sf.space_id                
JOIN
    features f ON sf.feature_id = f.id    
JOIN
    (
        SELECT DISTINCT ON (s.id) s.id, space_statuses.name
        FROM spaces AS s
        JOIN space_statuses ON s.status_id = space_statuses.id
        ORDER BY s.id
    ) AS space_statuses ON s.id = space_statuses.id  
GROUP BY s.id, space_statuses.name;

-- name: GetSpacePrices :many
SELECT
    s.id,    
    tpp.rate
FROM
    spaces AS s
JOIN
    pricing_groups pg ON s.group_id = pg.id
JOIN
    time_pricing_policy tpp ON tpp.group_id = pg.id   
WHERE
  tpp.day_of_week = $1 AnD tpp.hour = $2     
GROUP BY s.id, tpp.rate;

-- name: UpdateSpaceStatus :one
UPDATE spaces
  set status_id = $3
WHERE spaces.id = $1 AND spaces.status_id = $2
RETURNING id;

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

-- name: UpdateReservationStatus :exec
UPDATE space_reservations
  set status_id = $2
WHERE space_reservations.id = $1;

-- name: GetReservationStatuses :many
SELECT *
FROM reservation_statuses;

-- name: GetActiveReservations :many
SELECT id, time_from, time_to, car_id, status_id, space_id
FROM space_reservations as r
WHERE status_id == 1 OR status_id == 2;

-- name: GetReservationHistory :many
SELECT *
FROM space_reservations as r
WHERE status_id != 1 AND status_id != 2;

-- name: GetSpaceStatuses :many
SELECT *
FROM space_statuses;

-- name: DeleteSpaceFeatures :exec
DELETE FROM space_features
WHERE space_id = $1;

-- name: AssignSpaceFeature :one
INSERT INTO space_features (space_id, feature_id, is_required) 
VALUES ($1, $2, $3) RETURNING id;

