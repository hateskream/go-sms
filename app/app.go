package app

import (
	"context"
	"errors"
	"log"
	"space-management-system/services/db/db"
	"sync"

	"github.com/jackc/pgx/v5/pgtype"
	"modernc.org/libc/time"
)

type ActiveReservation struct {
	db.GetActiveReservationsRow
	Status string
}
type SpaceWithPrice struct {
	db.GetSpacesRow
	Price float32 `json:"price"`
}

type SpaceFeature struct {
	Name       string
	IsRequired bool
}

type StorageManager interface {
	GetSpaces(ctx context.Context) ([]db.GetSpacesRow, error)
	AddSpace(ctx context.Context, arg db.AddSpaceParams) (int32, error)
	GetFeatures(ctx context.Context) ([]db.Feature, error)
	GetSpacePrices(ctx context.Context, arg db.GetSpacePricesParams) ([]db.GetSpacePricesRow, error)
	GetSpaceStatuses(ctx context.Context) ([]db.SpaceStatus, error)
	UpdateSpaceStatus(ctx context.Context, arg db.UpdateSpaceStatusParams) (int32, error)
	DeleteSpaceFeatures(ctx context.Context, spaceID pgtype.Int4) error
	AssignSpaceFeature(ctx context.Context, arg db.AssignSpaceFeatureParams) (int32, error)

	AddFeature(ctx context.Context, name string) (int32, error)
	DeleteFeature(ctx context.Context, id int32) (int32, error)
	AddCardNumber(ctx context.Context, number string) (int32, error)
	GetCardByNumber(ctx context.Context, number string) (int32, error)

	AddPricingGroup(ctx context.Context, name string) (int32, error)
	GeneratTimePricingPolicy(ctx context.Context, arg db.GeneratTimePricingPolicyParams) error
	DeletePricingGroup(ctx context.Context, id int32) (int32, error)

	UpdatePricingPolicy(ctx context.Context, arg db.UpdatePricingPolicyParams) error
	UpdatePricingGroups(ctx context.Context, arg db.UpdatePricingGroupsParams) error

	AddReservation(ctx context.Context, arg db.AddReservationParams) (int32, error)
	GetReservationStatuses(ctx context.Context) ([]db.ReservationStatus, error)
	GetActiveReservations(ctx context.Context) ([]db.GetActiveReservationsRow, error)
	UpdateReservationStatus(ctx context.Context, arg db.UpdateReservationStatusParams) error
}

type HardwareManager interface {
	SetLocker(id int, up bool) error
}

type SpaceManager interface {
	SetApp(*App) error
	GetFeatures() ([]db.Feature, error)
	GetSpaces() ([]SpaceWithPrice, error)
	UpdateSpaceStatuses() error //updates space statuses in memory storage if needed
	AddSpace(name string, physical_id int32, groupId int32, status string) (int32, error)
	UpdateSpace(s *SpaceWithPrice) (int32, error)
	DeleteSpace(id int32) error
	UpdateSpacePrices() error //updates space prices in memory storage if needed
	UpdateSpaceStatus(space_id int32, curStatus string, newStatus string) error
	AssignSpaceFeatures(space_id int32, features []SpaceFeature) error
	AddFeature(f *db.Feature) (int32, error)
	RenameFeature(id int32, name string) error
	DeleteFeature(id int32) error
	AddPricingGroup(name string) (int32, error)
	GetReservationStatuses(ctx context.Context) ([]db.ReservationStatus, error)
}

type ReservationsManager interface {
	SetApp(*App) error
	AddNewReservation(car_number string, space_id int32, rservation_fee float32, time_to time.Time) (int32, error)
	ChangeReservationStatus(id int32, status string) error
	GetReservationHistory() ([]string, error)
	GetActiveReservations() ([]string, error)
	UpdateReservationStatuses() error //updates reservation statuses in memory storage if needed
	GetCarId(car_number string) (int32, error)
}

type App struct {
	Storage      StorageManager
	Spaces       SpaceManager
	Reservations ReservationsManager
	Hardware     HardwareManager
}

var appInstance *App
var lock = &sync.Mutex{}

func InitializeApp() *App {
	if appInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if appInstance == nil {
			log.Println("App initialized")
			appInstance = &App{}
		}
	}
	return appInstance
}

func (app *App) SetSpaces(sm SpaceManager) {
	app.Spaces = sm
	sm.SetApp(app)
}

func (app *App) SetReservations(rm ReservationsManager) {
	app.Reservations = rm
	rm.SetApp(app)
}

func GetStorage() (StorageManager, error) {
	if appInstance == nil {
		return nil, errors.New("App is not initialized yet")
	}
	return appInstance.Storage, nil
}

func GetSpacesManager() (SpaceManager, error) {
	if appInstance == nil {
		return nil, errors.New("App is not initialized yet")
	}
	return appInstance.Spaces, nil
}

func GetReservationManager() (ReservationsManager, error) {
	if appInstance == nil {
		return nil, errors.New("App is not initialized yet")
	}
	return appInstance.Reservations, nil
}
