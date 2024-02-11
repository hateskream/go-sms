package app

import (
	"context"
	"net/http"
	"space-management-system/services/db/db"
)

type HandlersInterface interface {
	GetSpaces(w http.ResponseWriter, r *http.Request)
	GetFeatures(w http.ResponseWriter, r *http.Request)
	ReserveSpace(w http.ResponseWriter, r *http.Request)
	UpdateReservationStatus(w http.ResponseWriter, r *http.Request)

	CarArrival(w http.ResponseWriter, r *http.Request)
	CarDeparture(w http.ResponseWriter, r *http.Request)

	GetSpacesInfo(w http.ResponseWriter, r *http.Request)
	AddSpace(w http.ResponseWriter, r *http.Request)
	DeleteSpace(w http.ResponseWriter, r *http.Request)
	UpdateSpace(w http.ResponseWriter, r *http.Request)
	AddFeature(w http.ResponseWriter, r *http.Request)
	DeleteFeature(w http.ResponseWriter, r *http.Request)
	UpdateFeature(w http.ResponseWriter, r *http.Request)
	AddPricingPolicy(w http.ResponseWriter, r *http.Request)
	DeletePricingPolicy(w http.ResponseWriter, r *http.Request)
	UpdatePricingPolicy(w http.ResponseWriter, r *http.Request)
	AddPricingGroup(w http.ResponseWriter, r *http.Request)
	DeletePricingGroup(w http.ResponseWriter, r *http.Request)
	UpdatePricingGroup(w http.ResponseWriter, r *http.Request)
	SetLocker(w http.ResponseWriter, r *http.Request)
	GetActiveReservations(w http.ResponseWriter, r *http.Request)
	GetReservationHistory(w http.ResponseWriter, r *http.Request)
	GetSpaceOccupancy(w http.ResponseWriter, r *http.Request)
}

type StorageInterface interface {
	GetAllSpaces(ctx context.Context) ([]db.GetAllSpacesRow, error)
	AddSpace(ctx context.Context, arg db.AddSpaceParams) (db.AddSpaceRow, error)
	UpdateSpace(ctx context.Context, arg db.UpdateSpaceParams) error
	DeleteSpace(ctx context.Context, id int32) (int32, error)
	GetSpacesByFeatureList(ctx context.Context, arg db.GetSpacesByFeatureListParams) ([]db.GetSpacesByFeatureListRow, error)
	GetFeatures(ctx context.Context) ([]db.Feature, error)
	AddFeature(ctx context.Context, name string) (int32, error)
	DeleteFeature(ctx context.Context, id int32) (int32, error)
	UpdateFeature(ctx context.Context, arg db.UpdateFeatureParams) error
	AddCardNumber(ctx context.Context, number string) (int32, error)
	GetCardByNumber(ctx context.Context, number string) (int32, error)
	UpdateSpaceStatus(ctx context.Context, arg db.UpdateSpaceStatusParams) error
	AddReservation(ctx context.Context, arg db.AddReservationParams) (int32, error)
	UpdateReservationState(ctx context.Context, arg db.UpdateReservationStateParams) error
	GetSpaceStatuses(ctx context.Context) ([]db.SpaceStatus, error)
}

type HardwareInterface interface {
	SetLocker(id int) error
}

type MemoryStorageInterface interface {
	setState()
}

type App struct {
	Storage  StorageInterface
	Hardware HardwareInterface
	Handlers HandlersInterface
	Memory   MemoryStorageInterface
}
