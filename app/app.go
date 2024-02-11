package app

import (
	"context"
	"errors"
	"log"
	"space-management-system/services/db/db"
	"sync"
)

type SpacesInterface interface {
}

type StorageManager interface {
	GetAllSpaces(ctx context.Context) ([]db.GetAllSpacesRow, error)
	GetSpacesByFeatureList(ctx context.Context, arg db.GetSpacesByFeatureListParams) ([]db.GetSpacesByFeatureListRow, error)
	GetFeatures(ctx context.Context) ([]db.Feature, error)
	AddFeature(ctx context.Context, name string) (int32, error)
	DeleteFeature(ctx context.Context, id int32) (int32, error)
	AddCardNumber(ctx context.Context, number string) (int32, error)
	GetCardByNumber(ctx context.Context, number string) (int32, error)
	UpdateSpaceStatus(ctx context.Context, arg db.UpdateSpaceStatusParams) error
	AddReservation(ctx context.Context, arg db.AddReservationParams) (int32, error)
	UpdateReservationState(ctx context.Context, arg db.UpdateReservationStateParams) error
	GetSpaceStatuses(ctx context.Context) ([]db.SpaceStatus, error)
}

type HardwareManager interface {
	SetLocker(id int) error
}

type SpaceManager interface {
	SetApp(*App)
	GetFeatures() ([]db.Feature, error)
	GetSpaces() ([]db.Space, error)
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
