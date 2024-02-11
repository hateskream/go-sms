// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: spaces.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addCardNumber = `-- name: AddCardNumber :one
INSERT INTO cars (number)
VALUES ($1::VARCHAR) RETURNING id
`

func (q *Queries) AddCardNumber(ctx context.Context, number string) (int32, error) {
	row := q.db.QueryRow(ctx, addCardNumber, number)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const addFeature = `-- name: AddFeature :one
INSERT INTO features (name)
VALUES ($1) RETURNING id
`

func (q *Queries) AddFeature(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRow(ctx, addFeature, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const addPricingGroup = `-- name: AddPricingGroup :one
INSERT INTO pricing_groups (name)
VALUES ($1) RETURNING id
`

func (q *Queries) AddPricingGroup(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRow(ctx, addPricingGroup, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const addReservation = `-- name: AddReservation :one
INSERT INTO space_reservations (time_from, time_to, car_id, reservation_fee, space_id, status_id)
VALUES ($1,$2,$3,$4,$5,$6) RETURNING id
`

type AddReservationParams struct {
	TimeFrom       pgtype.Timestamp `json:"time_from"`
	TimeTo         pgtype.Timestamp `json:"time_to"`
	CarID          pgtype.Int4      `json:"car_id"`
	ReservationFee float32          `json:"reservation_fee"`
	SpaceID        pgtype.Int4      `json:"space_id"`
	StatusID       pgtype.Int4      `json:"status_id"`
}

func (q *Queries) AddReservation(ctx context.Context, arg AddReservationParams) (int32, error) {
	row := q.db.QueryRow(ctx, addReservation,
		arg.TimeFrom,
		arg.TimeTo,
		arg.CarID,
		arg.ReservationFee,
		arg.SpaceID,
		arg.StatusID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const addSpace = `-- name: AddSpace :one
INSERT INTO spaces (name, physical_id, group_id, status_id)
VALUES ($1,$2,$3,$4)
RETURNING  name, physical_id, group_id, status_id
`

type AddSpaceParams struct {
	Name       string      `json:"name"`
	PhysicalID pgtype.Int4 `json:"physical_id"`
	GroupID    pgtype.Int4 `json:"group_id"`
	StatusID   pgtype.Int4 `json:"status_id"`
}

type AddSpaceRow struct {
	Name       string      `json:"name"`
	PhysicalID pgtype.Int4 `json:"physical_id"`
	GroupID    pgtype.Int4 `json:"group_id"`
	StatusID   pgtype.Int4 `json:"status_id"`
}

func (q *Queries) AddSpace(ctx context.Context, arg AddSpaceParams) (AddSpaceRow, error) {
	row := q.db.QueryRow(ctx, addSpace,
		arg.Name,
		arg.PhysicalID,
		arg.GroupID,
		arg.StatusID,
	)
	var i AddSpaceRow
	err := row.Scan(
		&i.Name,
		&i.PhysicalID,
		&i.GroupID,
		&i.StatusID,
	)
	return i, err
}

const addSpaceFeature = `-- name: AddSpaceFeature :exec

INSERT INTO space_features (space_id,Feature_id)
VALUES ($1,$2)
`

type AddSpaceFeatureParams struct {
	SpaceID   pgtype.Int4 `json:"space_id"`
	FeatureID pgtype.Int4 `json:"feature_id"`
}

//
func (q *Queries) AddSpaceFeature(ctx context.Context, arg AddSpaceFeatureParams) error {
	_, err := q.db.Exec(ctx, addSpaceFeature, arg.SpaceID, arg.FeatureID)
	return err
}

const cleanTimePricingPolicy = `-- name: CleanTimePricingPolicy :exec
DELETE FROM time_pricing_policy
WHERE group_id = $1
`

func (q *Queries) CleanTimePricingPolicy(ctx context.Context, groupID int32) error {
	_, err := q.db.Exec(ctx, cleanTimePricingPolicy, groupID)
	return err
}

const deleteFeature = `-- name: DeleteFeature :one
DELETE FROM features
WHERE id = $1 RETURNING id
`

func (q *Queries) DeleteFeature(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteFeature, id)
	err := row.Scan(&id)
	return id, err
}

const deletePricingGroup = `-- name: DeletePricingGroup :one
DELETE FROM pricing_groups
WHERE id = $1 RETURNING id
`

func (q *Queries) DeletePricingGroup(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deletePricingGroup, id)
	err := row.Scan(&id)
	return id, err
}

const deleteSpace = `-- name: DeleteSpace :one
DELETE FROM spaces
WHERE id = $1 RETURNING id
`

func (q *Queries) DeleteSpace(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteSpace, id)
	err := row.Scan(&id)
	return id, err
}

const generatTimePricingPolicy = `-- name: GeneratTimePricingPolicy :exec
INSERT INTO time_pricing_policy (rate, hour, day_of_week, group_id)
SELECT $2, hour, day_of_week, $1
FROM generate_series(1, 8) AS day_of_week
CROSS JOIN generate_series(1, 24) AS hour
`

type GeneratTimePricingPolicyParams struct {
	GroupID int32   `json:"group_id"`
	Rate    float32 `json:"rate"`
}

func (q *Queries) GeneratTimePricingPolicy(ctx context.Context, arg GeneratTimePricingPolicyParams) error {
	_, err := q.db.Exec(ctx, generatTimePricingPolicy, arg.GroupID, arg.Rate)
	return err
}

const getAllSpaces = `-- name: GetAllSpaces :many
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
GROUP BY s.id, space_statuses.name
`

type GetAllSpacesRow struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Features string `json:"features"`
}

func (q *Queries) GetAllSpaces(ctx context.Context) ([]GetAllSpacesRow, error) {
	rows, err := q.db.Query(ctx, getAllSpaces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllSpacesRow
	for rows.Next() {
		var i GetAllSpacesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Status,
			&i.Features,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCardByNumber = `-- name: GetCardByNumber :one
SELECT cars.id
FROM cars
WHERE cars.number = $1::VARCHAR
`

func (q *Queries) GetCardByNumber(ctx context.Context, number string) (int32, error) {
	row := q.db.QueryRow(ctx, getCardByNumber, number)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getFeatures = `-- name: GetFeatures :many
SELECT id, name
FROM features
ORDER BY id
`

func (q *Queries) GetFeatures(ctx context.Context) ([]Feature, error) {
	rows, err := q.db.Query(ctx, getFeatures)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feature
	for rows.Next() {
		var i Feature
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSpaceStatuses = `-- name: GetSpaceStatuses :many
Select id, name
From space_statuses
`

func (q *Queries) GetSpaceStatuses(ctx context.Context) ([]SpaceStatus, error) {
	rows, err := q.db.Query(ctx, getSpaceStatuses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SpaceStatus
	for rows.Next() {
		var i SpaceStatus
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSpaces = `-- name: GetSpaces :many
SELECT id, name, physical_id, group_id, status_id, has_camera
FROM spaces
ORDER BY id
`

func (q *Queries) GetSpaces(ctx context.Context) ([]Space, error) {
	rows, err := q.db.Query(ctx, getSpaces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Space
	for rows.Next() {
		var i Space
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.PhysicalID,
			&i.GroupID,
			&i.StatusID,
			&i.HasCamera,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSpacesByFeatureList = `-- name: GetSpacesByFeatureList :many
SELECT s.id, s.name
FROM 
  spaces s
JOIN 
  space_features sf ON s.id = sf.space_id
WHERE sf.feature_id = ANY($1::int[]) AND s.status_id = 1
GROUP BY s.id
HAVING  COUNT(DISTINCT sf.feature_id) = $2::int
`

type GetSpacesByFeatureListParams struct {
	FeatureList  []int32 `json:"feature_list"`
	FeatureCount int32   `json:"feature_count"`
}

type GetSpacesByFeatureListRow struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) GetSpacesByFeatureList(ctx context.Context, arg GetSpacesByFeatureListParams) ([]GetSpacesByFeatureListRow, error) {
	rows, err := q.db.Query(ctx, getSpacesByFeatureList, arg.FeatureList, arg.FeatureCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSpacesByFeatureListRow
	for rows.Next() {
		var i GetSpacesByFeatureListRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFeature = `-- name: UpdateFeature :exec
UPDATE features
  set name = $2  
WHERE id = $1
`

type UpdateFeatureParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateFeature(ctx context.Context, arg UpdateFeatureParams) error {
	_, err := q.db.Exec(ctx, updateFeature, arg.ID, arg.Name)
	return err
}

const updatePricingGroups = `-- name: UpdatePricingGroups :exec
UPDATE pricing_groups
  set name = $2
  WHERE id = $1
`

type UpdatePricingGroupsParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdatePricingGroups(ctx context.Context, arg UpdatePricingGroupsParams) error {
	_, err := q.db.Exec(ctx, updatePricingGroups, arg.ID, arg.Name)
	return err
}

const updatePricingPolicy = `-- name: UpdatePricingPolicy :exec
UPDATE time_pricing_policy
  set rate = $2
  WHERE ID = $1
`

type UpdatePricingPolicyParams struct {
	ID   int32   `json:"id"`
	Rate float32 `json:"rate"`
}

func (q *Queries) UpdatePricingPolicy(ctx context.Context, arg UpdatePricingPolicyParams) error {
	_, err := q.db.Exec(ctx, updatePricingPolicy, arg.ID, arg.Rate)
	return err
}

const updateReservationState = `-- name: UpdateReservationState :exec
UPDATE space_reservations
  set status_id = $2
WHERE space_reservations.id = $1
`

type UpdateReservationStateParams struct {
	ID       int32       `json:"id"`
	StatusID pgtype.Int4 `json:"status_id"`
}

func (q *Queries) UpdateReservationState(ctx context.Context, arg UpdateReservationStateParams) error {
	_, err := q.db.Exec(ctx, updateReservationState, arg.ID, arg.StatusID)
	return err
}

const updateSpace = `-- name: UpdateSpace :exec
UPDATE spaces
  set name = $2,
  physical_id = $3,
  group_id = $4,
  status_id = $5
WHERE id = $1
`

type UpdateSpaceParams struct {
	ID         int32       `json:"id"`
	Name       string      `json:"name"`
	PhysicalID pgtype.Int4 `json:"physical_id"`
	GroupID    pgtype.Int4 `json:"group_id"`
	StatusID   pgtype.Int4 `json:"status_id"`
}

func (q *Queries) UpdateSpace(ctx context.Context, arg UpdateSpaceParams) error {
	_, err := q.db.Exec(ctx, updateSpace,
		arg.ID,
		arg.Name,
		arg.PhysicalID,
		arg.GroupID,
		arg.StatusID,
	)
	return err
}

const updateSpaceStatus = `-- name: UpdateSpaceStatus :exec
UPDATE spaces
  set status_id = $3
WHERE spaces.id = $1 AND spaces.status_id = $2
`

type UpdateSpaceStatusParams struct {
	ID         int32       `json:"id"`
	StatusID   pgtype.Int4 `json:"status_id"`
	StatusID_2 pgtype.Int4 `json:"status_id_2"`
}

func (q *Queries) UpdateSpaceStatus(ctx context.Context, arg UpdateSpaceStatusParams) error {
	_, err := q.db.Exec(ctx, updateSpaceStatus, arg.ID, arg.StatusID, arg.StatusID_2)
	return err
}
