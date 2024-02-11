package spaces

import (
	"context"
	"log"
	"space-management-system/app"
	"space-management-system/services/db/db"
)

type Space struct {
	db.Space
	status   string
	features []string
	price    float32
}

type SpacesManager struct {
	App                *app.App
	onUpdate           func(*[]Space)
	spaces             []Space
	price_updated_hour int
	features           []db.Feature
	space_statuses     []db.SpaceStatus
}

func (sm *SpacesManager) SetApp(app *app.App) {
	sm.App = app
}

func (sm *SpacesManager) GetSpaces() ([]Space, error) { //get all spaces with features and prices
	if len(sm.spaces) > 0 {
		sm.UpdatePrices()
		return sm.spaces, nil
	}

	spaces, err := sm.App.Storage.GetAllSpaces(context.Background())

	if err != nil {
		return nil, err
	}

}

func (sm *SpacesManager) GetSpacesFromStorage() error { //get all spaces with features and prices from db

	sm.onUpdate(&sm.spaces)
	return nil
}

func (sm *SpacesManager) SetSpaceStatus(current string, new string) error { //change space status (current status for validation)
	return nil
}

func (sm *SpacesManager) UpdatePrices() error {
	sm.onUpdate(&sm.spaces)
	return nil
}

func (sm *SpacesManager) GetSpacePrices() error { //update prices for all spaces from db
	sm.onUpdate(&sm.spaces)
	return nil
}

func (sm *SpacesManager) AddSpace(s *Space) (int32, error) {

	sm.onUpdate(&sm.spaces)
	return 0, nil
}

func (sm *SpacesManager) UpdateSpace(s *Space) (int32, error) {

	sm.onUpdate(&sm.spaces)
	return 0, nil
}

func (sm *SpacesManager) DeleteSpace(id int32) error {

	sm.onUpdate(&sm.spaces)
	return nil
}

func (sm *SpacesManager) UpdateSpaceFeatures([]string) error {

	sm.onUpdate(&sm.spaces)
	return nil
}

func (sm *SpacesManager) GetFeaturesFromStorage() ([]db.Feature, error) {
	return sm.App.Storage.GetFeatures(context.Background())
}

func (sm *SpacesManager) GetFeatures() ([]db.Feature, error) { //Features CRUD
	if len(sm.features) > 0 {
		return sm.features, nil
	}
	log.Println("tut")
	features, err := sm.GetFeaturesFromStorage()
	if err != nil {
		return nil, err
	}

	sm.features = features
	return features, nil
}

func (sm *SpacesManager) AddFeature(f *db.Feature) (int32, error) {
	return 0, nil
}

func (sm *SpacesManager) RenameFeature(id int32, name string) error {
	return nil
}

func (sm *SpacesManager) DeleteFeature(id int32) error {
	return nil
}

// get pricing group
// add pricing group + timing policy for it (192 inserts)
// rename pricing group
// delete pricing group
// update pricing policy (pricing_group_id and vaiable length array or days with 24records for each day)
