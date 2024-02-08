package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func (h *Handlers) GetAllSpaces(w http.ResponseWriter, r *http.Request) {
	//get all spaces (id, name, feature_list, price) and return it to client
}

func (h *Handlers) FindSpaces(w http.ResponseWriter, r *http.Request) {
	//parse Request for params (required features)
	//get spaces by params
	//send response with the spaces
}

func (h *Handlers) GetFeatures(w http.ResponseWriter, r *http.Request) {
	features, err := h.Storage.GetFeatures(context.Background())
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	jsonData, err := json.Marshal(features)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
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
