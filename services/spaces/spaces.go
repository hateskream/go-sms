package spaces

import (
	"context"
	"errors"
	"log"
	"space-management-system/app"
	"space-management-system/services/db/db"

	"github.com/jackc/pgx/v5/pgtype"
)
=

type SpacesManager struct {
	App                *app.App
	onUpdate           func(*[]app.SpaceWithPrice)
	spaces             []app.SpaceWithPrice
	price_updated_hour int
	features           []db.Feature
	space_statuses     []db.SpaceStatus
}

func (sm *SpacesManager) SetApp(app *app.App) {
	sm.App = app
	sm.price_updated_hour = -1
	sm.GetSpaces()
}

func (sm *SpacesManager) spacesUpdated() {
	sm.UpdateSpacePrices()
	log.Println("spaces updated", sm.spaces)
	if sm.onUpdate != nil {
		sm.onUpdate(&sm.spaces)
	}
}

func (sm *SpacesManager) GetSpaces() ([]app.SpaceWithPrice, error) { //get all spaces with features and statuses
	if len(sm.spaces) > 0 {
		sm.UpdateSpacePrices()
		return sm.spaces, nil
	}

	spaces, err := sm.App.Storage.GetSpaces(context.Background())
	if err != nil {
		return nil, err
	}
	sm.spaces = nil
	for _, space := range spaces {
		sm.spaces = append(sm.spaces, app.SpaceWithPrice{
			GetSpacesRow: space,
		})
	}
	sm.spacesUpdated()
	return sm.spaces, nil
}

func (sm *SpacesManager) UpdateSpaceStatuses() error {
	if len(sm.space_statuses) == 0 {
		statuses, err := sm.App.Storage.GetSpaceStatuses(context.Background())
		if err != nil {
			return err
		}
		sm.space_statuses = statuses
	}
	return nil
}

func (sm *SpacesManager) UpdateSpaceStatus(space_id int32, curStatus string, newStatus string) error { //change space status (current status for validation)
	err := sm.UpdateSpaceStatuses()
	if err != nil {
		return err
	}

	params := db.UpdateSpaceStatusParams{ID: space_id}
	log.Println("sm statuses", sm.space_statuses)
	for _, status := range sm.space_statuses {
		if status.Name == curStatus {
			params.StatusID = pgtype.Int4{Int32: status.ID, Valid: true}
		}
		if status.Name == newStatus {
			params.StatusID_2 = pgtype.Int4{Int32: status.ID, Valid: true}
		}
	}

	id, err := sm.App.Storage.UpdateSpaceStatus(context.Background(), params)

	if id == 0 {
		return errors.New("Requested space is not open for reservation")
	}
	if err != nil {
		return err
	}

	for i, _ := range sm.spaces {
		if sm.spaces[i].ID == space_id {
			sm.spaces[i].Status = newStatus
			break
		}
	}
	sm.spacesUpdated()
	return nil
}

func (sm *SpacesManager) UpdateSpacePrices() error {
	// currentTime := time.Now().UTC()
	// currentHour := currentTime.Hour()
	currentHour := 0
	if currentHour != sm.price_updated_hour {
		// currentDay := int(currentTime.Weekday())
		currentDay := 0
		params := db.GetSpacePricesParams{
			DayOfWeek: int16(currentDay),
			Hour:      int16(currentHour),
		}
		prices, err := sm.App.Storage.GetSpacePrices(context.Background(), params)
		if err != nil {
			return err
		}

		pricesMap := make(map[int32]float32)
		for _, row := range prices {
			pricesMap[row.ID] = row.Rate
		}
		for i, space := range sm.spaces {
			sm.spaces[i].Price = pricesMap[space.ID]
		}
		sm.price_updated_hour = currentHour
	}
	return nil
}

func (sm *SpacesManager) AddSpace(s *app.SpaceWithPrice) (int32, error) {

	return 0, nil
}

func (sm *SpacesManager) UpdateSpace(s *app.SpaceWithPrice) (int32, error) {

	return 0, nil
}

func (sm *SpacesManager) DeleteSpace(id int32) error {

	return nil
}

func (sm *SpacesManager) UpdateSpaceFeatures(space_id int32, features app.SpaceFeature) error {
	ctx := context.Background()
	space := pgtype.Int4{Int32: space_id, Valid: true}
	err := sm.App.Storage.DeleteSpaceFeatures(ctx, space)
	if err != nil {
		return err
	}
	sm.GetFeatures()
	featureMap := make(map[string]int32)
	for _, feature := range sm.features {
		featureMap[feature.Name] = feature.ID
	}
	for _, feature := range features {
		featureID := int32(featureMap[feature.Name])
		params := db.AssignSpaceFeatureParams{
			SpaceID:    space,
			FeatureID:  pgtype.Int4{Int32: featureID, Valid: true},
			IsRequired: feature.IsRequired,
		}
		id, err := sm.App.Storage.AssignSpaceFeature(ctx, params)
		if id == 0 || err != nil {
			return errors.New("error assigning new feature")
		}
	}

	return nil
}

func (sm *SpacesManager) GetFeatures() ([]db.Feature, error) { //Features CRUD
	if len(sm.features) > 0 {
		return sm.features, nil
	}
	features, err := sm.App.Storage.GetFeatures(context.Background())
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
