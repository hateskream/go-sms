package spaces

import (
	"context"
	"errors"
	"fmt"
	"log"
	"space-management-system/app"
	"space-management-system/services/db/db"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

type SpacesManager struct {
	App                *app.App
	onUpdate           func(*[]app.SpaceWithPrice)
	spaces             []app.SpaceWithPrice
	price_updated_hour int
	features           []db.Feature
	space_statuses     []db.SpaceStatus
}

func (sm *SpacesManager) SetApp(app *app.App) error {
	sm.App = app
	sm.price_updated_hour = -1
	sm.GetSpaces()
	return nil
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
		stringValue, ok := space.Features.(string)
		if ok {
			space.Features = strings.Split(string(stringValue), ",")
		}
		stringValue, ok = space.RequiredFeatures.(string)
		if ok {
			space.RequiredFeatures = strings.Split(string(stringValue), ",")
		}
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
	for _, status := range sm.space_statuses {
		if status.Name == curStatus {
			params.StatusID = pgtype.Int4{Int32: status.ID, Valid: true}
		}
		if status.Name == newStatus {
			params.StatusID_2 = pgtype.Int4{Int32: status.ID, Valid: true}
		}
	}
	log.Println("")
	id, err := sm.App.Storage.UpdateSpaceStatus(context.Background(), params)

	if id == 0 {
		return fmt.Errorf("the space is not %s", curStatus)
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

func (sm *SpacesManager) AddSpace(name string, physicalId int32, groupId int32, status string) (int32, error) {
	sm.UpdateSpaceStatuses()
	if groupId == 0 || physicalId == 0 {
		status = "disabled"
	}
	if groupId == 0 {
		groupId = 1
	}

	statusId := int32(0)
	for _, s := range sm.space_statuses {
		if s.Name == status {
			statusId = s.ID
		}

		if statusId == 0 && s.Name == "disabled" {
			statusId = s.ID
		}
	}

	params := db.AddSpaceParams{
		Name:       name,
		PhysicalID: pgtype.Int4{Int32: int32(physicalId), Valid: true},
		GroupID:    pgtype.Int4{Int32: int32(groupId), Valid: true},
		StatusID:   pgtype.Int4{Int32: int32(statusId), Valid: true},
	}

	added, err := sm.App.Storage.AddSpace(context.Background(), params)
	if err != nil {
		return 0, err
	}

	sm.spaces = append(sm.spaces, app.SpaceWithPrice{
		GetSpacesRow: db.GetSpacesRow{
			ID:               added,
			Name:             name,
			Status:           status,
			RequiredFeatures: nil,
			Features:         nil,
		},
	})
	sm.spacesUpdated()

	return added, nil
}

func (sm *SpacesManager) UpdateSpace(s *app.SpaceWithPrice) (int32, error) {

	return 0, nil
}

func (sm *SpacesManager) DeleteSpace(id int32) error {

	return nil
}

func (sm *SpacesManager) AssignSpaceFeatures(space_id int32, features []app.SpaceFeature) error {
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

	var featureList []string
	var requiredFeatureList []string

	for _, feature := range features {
		featureID := int32(featureMap[feature.Name])
		params := db.AssignSpaceFeatureParams{
			SpaceID:    space,
			FeatureID:  pgtype.Int4{Int32: featureID, Valid: true},
			IsRequired: feature.IsRequired,
		}
		featureList = append(featureList, feature.Name)
		if feature.IsRequired == true {
			requiredFeatureList = append(requiredFeatureList, feature.Name)
		}
		id, err := sm.App.Storage.AssignSpaceFeature(ctx, params)
		if id == 0 || err != nil {
			return errors.New("error assigning new feature")
		}
	}

	for i := range sm.spaces {
		if sm.spaces[i].ID == space_id {
			sm.spaces[i].Features = featureList
			sm.spaces[i].RequiredFeatures = requiredFeatureList
			break
		}
	}
	sm.spacesUpdated()

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

func (sm *SpacesManager) AddPricingGroup(name string) (int32, error) {

	group_id, err := sm.App.Storage.AddPricingGroup(context.Background(), name)
	if err != nil {
		return 0, err
	}

	DEFAULT_RATE := 100
	params := db.GeneratTimePricingPolicyParams{
		GroupID: group_id,
		Rate:    float32(DEFAULT_RATE),
	}

	err = sm.App.Storage.GeneratTimePricingPolicy(context.Background(), params)
	if err != nil {
		return 0, err
	}
	return group_id, nil
}

// get pricing group
// update pricing policy (pricing_group_id and vaiable length array or days with 24records for each day)
