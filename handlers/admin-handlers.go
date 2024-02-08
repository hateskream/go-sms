package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetSpacesInfo(w http.ResponseWriter, r *http.Request) {
	//get space id, name, status, group and list of features + pagination
}

func (h *Handlers) AddSpace(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) DeleteSpace(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) UpdateSpace(w http.ResponseWriter, r *http.Request) {
	//name, features, status, pricing group
}

func (h *Handlers) AddFeature(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	if name == "" {
		http.Error(w, "name is required field", 422)
		return
	}
	added_id, err := h.Storage.AddFeature(context.Background(), name)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	data := map[string]int32{
		"id": added_id,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write(jsonData)
}

func (h *Handlers) DeleteFeature(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "featureID")
	idInt, _ := strconv.Atoi(id)
	deleted_id, err := h.Storage.DeleteFeature(context.Background(), int32(idInt))
	log.Println(deleted_id, err)
	if err != nil {
		errMsg := err.Error()
		http.Error(w, errMsg, 500)
		return
	}
	response := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "Feature successfully deleted.",
	}
	jsonData, _ := json.Marshal(response)
	w.Write(jsonData)
}

func (h *Handlers) UpdateFeature(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) AddPricingPolicy(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) DeletePricingPolicy(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) UpdatePricingPolicy(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) AddPricingGroup(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) DeletePricingGroup(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) UpdatePricingGroup(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) SetLocker(w http.ResponseWriter, r *http.Request) {
	//set locker position and space status + check for reservations
}

func (h *Handlers) GetActiveReservations(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetReservationHistory(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetSpaceOccupancy(w http.ResponseWriter, r *http.Request) {
	//by date
}
