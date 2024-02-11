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

func UpdateSpaceFeature(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	if id == "" {
		http.Error(w, "id is required field", 422)
		return
	}
	space_id, _ := strconv.Atoi(id)
	var featuresSlice []app.SpaceFeature
	featuresJson := r.Form.Get("features")
	if featuresJson != "" {
		features := json.Unmarshal([]byte(featuresJson))
	}
	sm, _ := app.GetSpacesManager()
	err := sm.UpdateSpaceFeatures(int32(space_id), featuresSlice)
}

// func AddSpace(w http.ResponseWriter, r *http.Request) {

// 	r.ParseForm()
// 	name := r.Form.Get("name")
// 	if name == "" {
// 		http.Error(w, "name is required field", 422)
// 		return
// 	}
// 	physical_id := r.Form.Get("physical_id")
// 	physical_idInt, _ := strconv.Atoi(physical_id)
// 	group_id := r.Form.Get("group_id")
// 	group_idInt, _ := strconv.Atoi(group_id)
// 	status_id := r.Form.Get("status_id")
// 	status_idInt, _ := strconv.Atoi(status_id)
// 	params := db.AddSpaceParams{
// 		Name:       name,
// 		PhysicalID: pgtype.Int4{Int32: int32(physical_idInt), Valid: true},
// 		GroupID:    pgtype.Int4{Int32: int32(group_idInt), Valid: true},
// 		StatusID:   pgtype.Int4{Int32: int32(status_idInt), Valid: true},
// 	}

// 	log.Printf("GetRequest: %s %s %s %s", name, physical_id, group_id, status_id)
// 	added, err := h.Storage.AddSpace(context.Background(), params)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	jsonData, err := json.Marshal(added)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	w.Write(jsonData)
// }

// func DeleteSpace(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "spaceID")
// 	idInt, _ := strconv.Atoi(id)
// 	deleted_id, err := h.Storage.DeleteSpace(context.Background(), int32(idInt))
// 	log.Println(deleted_id, err)
// 	if err != nil {
// 		errMsg := err.Error()
// 		http.Error(w, errMsg, 500)
// 		return
// 	}
// 	response := struct {
// 		Success bool   `json:"success"`
// 		Message string `json:"message"`
// 	}{
// 		Success: true,
// 		Message: "Feature successfully deleted.",
// 	}
// 	jsonData, _ := json.Marshal(response)
// 	w.Write(jsonData)
// }

// func UpdateSpace(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "spaceID")
// 	idInt, _ := strconv.Atoi(id)
// 	r.ParseForm()
// 	name := r.Form.Get("name")
// 	log.Printf("id: %s %s", id, name)
// 	if name == "" {
// 		http.Error(w, "name is required field", 422)
// 		return
// 	}
// 	physical_id := r.Form.Get("physical_id")
// 	physical_idInt, _ := strconv.Atoi(physical_id)
// 	group_id := r.Form.Get("group_id")
// 	group_idInt, _ := strconv.Atoi(group_id)
// 	status_id := r.Form.Get("status_id")
// 	status_idInt, _ := strconv.Atoi(status_id)

// 	params := db.UpdateSpaceParams{
// 		ID:         int32(idInt),
// 		Name:       name,
// 		PhysicalID: pgtype.Int4{Int32: int32(physical_idInt), Valid: true},
// 		GroupID:    pgtype.Int4{Int32: int32(group_idInt), Valid: true},
// 		StatusID:   pgtype.Int4{Int32: int32(status_idInt), Valid: true},
// 	}

// 	err := h.Storage.UpdateSpace(context.Background(), params)

// 	if err != nil {
// 		errMsg := err.Error()
// 		http.Error(w, errMsg, 500)
// 		return
// 	}
// 	response := struct {
// 		Success bool   `json:"success"`
// 		Message string `json:"message"`
// 	}{
// 		Success: true,
// 		Message: "Feature successfully updated.",
// 	}
// 	jsonData, _ := json.Marshal(response)
// 	w.Write(jsonData)
// }

// func AddFeature(w http.ResponseWriter, r *http.Request) {
// 	storage, _ := app.GetStorage()
// 	r.ParseForm()
// 	name := r.Form.Get("name")
// 	if name == "" {
// 		http.Error(w, "name is required field", 422)
// 		return
// 	}
// 	added_id, err := storage.AddFeature(context.Background(), name)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	data := map[string]int32{
// 		"id": added_id,
// 	}
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	w.Write(jsonData)
// }

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

// func UpdateFeature(w http.ResponseWriter, r *http.Request) {
// 	storage, _ := app.GetStorage()
// 	id := chi.URLParam(r, "featureID")
// 	idInt, _ := strconv.Atoi(id)

// 	r.ParseForm()
// 	name := r.Form.Get("name")

// 	params := db.UpdateFeatureParams{
// 		ID:   int32(idInt),
// 		Name: name,
// 	}
// 	err := storage.UpdateFeature(context.Background(), params)

// 	if err != nil {
// 		errMsg := err.Error()
// 		http.Error(w, errMsg, 500)
// 		return
// 	}

// 	response := struct {
// 		Success bool   `json:"success"`
// 		Message string `json:"message"`
// 	}{
// 		Success: true,
// 		Message: "Feature successfully updated.",
// 	}
// 	jsonData, _ := json.Marshal(response)
// 	w.Write(jsonData)

// }

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
