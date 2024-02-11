package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"space-management-system/app"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AddSpace(w http.ResponseWriter, r *http.Request) {

}

func DeleteSpace(w http.ResponseWriter, r *http.Request) {

}

func UpdateSpace(w http.ResponseWriter, r *http.Request) {
	//name, features, status, pricing group
}

func AddFeature(w http.ResponseWriter, r *http.Request) {
	storage, _ := app.GetStorage()
	r.ParseForm()
	name := r.Form.Get("name")
	if name == "" {
		http.Error(w, "name is required field", 422)
		return
	}
	added_id, err := storage.AddFeature(context.Background(), name)
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

func DeleteFeature(w http.ResponseWriter, r *http.Request) {
	storage, _ := app.GetStorage()
	id := chi.URLParam(r, "featureID")
	idInt, _ := strconv.Atoi(id)
	deleted_id, err := storage.DeleteFeature(context.Background(), int32(idInt))
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

func UpdateFeature(w http.ResponseWriter, r *http.Request) {

}

func AddPricingPolicy(w http.ResponseWriter, r *http.Request) {

}

func DeletePricingPolicy(w http.ResponseWriter, r *http.Request) {

}

func UpdatePricingPolicy(w http.ResponseWriter, r *http.Request) {

}

func AddPricingGroup(w http.ResponseWriter, r *http.Request) {

}

func DeletePricingGroup(w http.ResponseWriter, r *http.Request) {

}

func UpdatePricingGroup(w http.ResponseWriter, r *http.Request) {

}

func SetLocker(w http.ResponseWriter, r *http.Request) {
	//set locker position and space status + check for reservations
}

func GetActiveReservations(w http.ResponseWriter, r *http.Request) {

}

func GetReservationHistory(w http.ResponseWriter, r *http.Request) {

}

func GetSpaceOccupancy(w http.ResponseWriter, r *http.Request) {
	//by date
}
