package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"space-management-system/app"
	"space-management-system/services/db/db"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

func GetSpaces(w http.ResponseWriter, r *http.Request) {
	sm, _ := app.GetSpacesManager()
	spaces, err := sm.GetSpaces()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	jsonData, err := json.Marshal(spaces)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(jsonData)

	// if ok {
	// 	featureslist, _ := storage.GetFeatures(context.Background())
	// 	var featureIds []int32
	// 	for _, el := range featureslist {
	// 		if utils.ArrayIncludes(features, el.Name) {
	// 			featureIds = append(featureIds, el.ID)
	// 		}
	// 	}
	// 	if len(features) != len(featureIds) {
	// 		http.Error(w, http.StatusText(422), 422)
	// 		return
	// 	}

	// 	params := db.GetSpacesByFeatureListParams{
	// 		FeatureList:  featureIds,
	// 		FeatureCount: int32(len(featureIds)),
	// 	}
	// 	spaces, err = storage.GetSpacesByFeatureList(context.Background(), params)
	// } else {
	// 	spaces, err = storage.GetAllSpaces(context.Background())
	// }

}

func FindSpacesByFeatureList(w http.ResponseWriter, r *http.Request) {
	features := r.URL.Query().Get("features")
	w.Write([]byte(features))
	//parse Request for params (required features)
	//get spaces by params
	//send response with the spaces
}

func GetFeatures(w http.ResponseWriter, r *http.Request) {
	sm, _ := app.GetSpacesManager()
	features, err := sm.GetFeatures()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	jsonData, err := json.Marshal(features)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(jsonData)
}

func ReserveSpace(w http.ResponseWriter, r *http.Request) {
	storage, _ := app.GetStorage()
	r.ParseForm()
	space_id, err := strconv.Atoi(r.Form.Get("space"))
	car_number := r.Form.Get("car_number")

	if car_number == "" {
		http.Error(w, "car number is required parameter", 422)
	}
	if err != nil {
		http.Error(w, "incorrect space_id", 422)
	}
	params := db.UpdateSpaceStatusParams{
		ID:         int32(space_id),
		StatusID:   pgtype.Int4{Int32: 1, Valid: true},
		StatusID_2: pgtype.Int4{Int32: 2, Valid: true},
	}

	storage.UpdateSpaceStatus(context.Background(), params)
	//Find user car or create one
	//Add reservation record
}

func UpdateReservationStatus(w http.ResponseWriter, r *http.Request) {
	//Find reservation by id and update it status
}
