// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Car struct {
	ID     int32       `json:"id"`
	Number pgtype.Text `json:"number"`
}

type Feature struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type ReservationStatus struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Space struct {
	ID         int32       `json:"id"`
	Name       string      `json:"name"`
	PhysicalID pgtype.Int4 `json:"physical_id"`
	GroupID    pgtype.Int4 `json:"group_id"`
	StatusID   pgtype.Int4 `json:"status_id"`
	HasCamera  bool        `json:"has_camera"`
}

type SpaceFeature struct {
	ID        int32       `json:"id"`
	SpaceID   pgtype.Int4 `json:"space_id"`
	FeatureID pgtype.Int4 `json:"feature_id"`
}

type SpaceGroup struct {
	ID                  int32       `json:"id"`
	Name                string      `json:"name"`
	ParkingPrice        float32     `json:"parking_price"`
	BookingStaticPrice  float32     `json:"booking_static_price"`
	TimePricingPolicyID pgtype.Int4 `json:"time_pricing_policy_id"`
}

type SpaceOccupancy struct {
	ID         int32            `json:"id"`
	SpaceID    pgtype.Int4      `json:"space_id"`
	Timestamp  pgtype.Timestamp `json:"timestamp"`
	IsOccupied bool             `json:"is_occupied"`
	CarID      pgtype.Int4      `json:"car_id"`
}

type SpaceReservation struct {
	ID                  int32            `json:"id"`
	TimeFrom            pgtype.Timestamp `json:"time_from"`
	TimeTo              pgtype.Timestamp `json:"time_to"`
	CarID               pgtype.Int4      `json:"car_id"`
	ReservationFee      float32          `json:"reservation_fee"`
	StatusID            pgtype.Int4      `json:"status_id"`
	SpaceID             pgtype.Int4      `json:"space_id"`
	ParkingTimeFrom     pgtype.Timestamp `json:"parking_time_from"`
	ParkingTimeTo       pgtype.Timestamp `json:"parking_time_to"`
	ParkingFee          int32            `json:"parking_fee"`
	ParkingFeeBreakdown int32            `json:"parking_fee_breakdown"`
}

type SpaceStatus struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type TimePricingPolicy struct {
	ID        int32   `json:"id"`
	Rate      float32 `json:"rate"`
	Hour      int16   `json:"hour"`
	DayOfWeek int16   `json:"day_of_week"`
}