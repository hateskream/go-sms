package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"space-management-system/services/db/db"
	"space-management-system/utils"
)

func (h *Handlers) GetSpaces(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	features, ok := queryParams["features"]
	var spaces interface{}
	var err error
	if ok {
		featureslist, _ := h.Storage.GetFeatures(context.Background())
		var featureIds []int32
		for _, el := range featureslist {
			if utils.ArrayIncludes(features, el.Name) {
				featureIds = append(featureIds, el.ID)
			}
		}
		if len(features) != len(featureIds) {
			http.Error(w, http.StatusText(422), 422)
			return
		}

		params := db.GetSpacesByFeatureListParams{
			FeatureList:  featureIds,
			FeatureCount: int32(len(featureIds)),
		}
		spaces, err = h.Storage.GetSpacesByFeatureList(context.Background(), params)
	} else {
		spaces, err = h.Storage.GetAllSpaces(context.Background())
	}

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

}

func (h *Handlers) FindSpacesByFeatureList(w http.ResponseWriter, r *http.Request) {
	features := r.URL.Query().Get("features")
	w.Write([]byte(features))
	//parse Request for params (required features)
	//get spaces by params
	//send response with the spaces
}

func (h *Handlers) GetFeatures(w http.ResponseWriter, r *http.Request) {

	features, err := h.Storage.GetFeatures(context.Background())
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

func (h *Handlers) ReserveSpace(w http.ResponseWriter, r *http.Request) {
	//FindSpaces from params
	//Take the best(set it status to reserved)
	//Find user car or create one
	//Add reservation record
}

func (h *Handlers) UpdateReservationStatus(w http.ResponseWriter, r *http.Request) {
	//Find reservation by id and update it status
}
