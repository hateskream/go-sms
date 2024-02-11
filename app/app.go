package app

import (
	"context"
	"errors"
	"log"
	"space-management-system/services/db/db"
	"sync"

	"github.com/jackc/pgx/v5/pgtype"
)

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
	CleanTimePricingPolicy(ctx context.Context, id int32) error
	UpdatePricingPolicy(ctx context.Context, arg db.UpdatePricingPolicyParams) error
	UpdatePricingGroups(ctx context.Context, arg db.UpdatePricingGroupsParams) error
}

type HardwareManager interface {
	SetLocker(id int) error
}

type SpaceManager interface {
	SetApp(*App)
	GetFeatures() ([]db.Feature, error)
	GetSpaces() ([]SpaceWithPrice, error)
	UpdateSpacePrices() error
	UpdateSpaceStatus(space_id int32, curStatus string, newStatus string) error
	UpdateSpaceFeatures(space_id int32, features []SpaceFeature) error
}

type ReservationsManager interface {
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
