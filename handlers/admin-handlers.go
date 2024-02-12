package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"space-management-system/app"
	"space-management-system/services/db/db"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AddSpace(w http.ResponseWriter, r *http.Request) {
	sm, _ := app.GetSpacesManager()
	r.ParseForm()
	name := r.Form.Get("name")
	if name == "" {
		http.Error(w, "name is required field", 422)
		return
	}
	physical_id := r.Form.Get("physical_id")
	if physical_id == "" {
		physical_id = "0"
	}
	physical_idInt, _ := strconv.Atoi(physical_id)

	group_id := r.Form.Get("group_id")
	if group_id == "" {
		group_id = "0"
	}
	group_idInt, _ := strconv.Atoi(group_id)

	status := r.Form.Get("status")

	added, err := sm.AddSpace(name, int32(physical_idInt), int32(group_idInt), status)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	jsonData, err := json.Marshal(added)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(jsonData)
}

type updateFeatureRequest struct {
	ID       string
	Features []app.SpaceFeature
}

func UpdateSpaceFeatures(w http.ResponseWriter, r *http.Request) {
	var data updateFeatureRequest
	err := decodeJSONBody(w, r, &data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if data.ID == "" {
		http.Error(w, "id field is required", http.StatusUnprocessableEntity)
		return
	}
	space_id, _ := strconv.Atoi(data.ID)

	sm, _ := app.GetSpacesManager()

	err = sm.AssignSpaceFeatures(int32(space_id), data.Features)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(SuccessResponse("features were successfully updated"))

}

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

// func (h *Handlers) UpdateSpace(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "spaceID")
// 	if id == "" {
// 		errMsg := "Id is required"
// 		http.Error(w, errMsg, http.StatusUnprocessableEntity)
// 		return
// 	}
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

// func (h *Handlers) UpdateFeature(w http.ResponseWriter, r *http.Request) {

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

func UpdatePricingPolicy(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "featureID")
	// idInt, _ := strconv.Atoi(id)

	// r.ParseForm()
	// name := r.Form.Get("name")
}

func GetPricingPolicy(w http.ResponseWriter, r *http.Request) {

}

func AddPricingGroup(w http.ResponseWriter, r *http.Request) {
	sm, _ := app.GetSpacesManager()
	r.ParseForm()
	name := r.Form.Get("name")
	if name == "" {
		errMsg := "Name is required"
		http.Error(w, errMsg, http.StatusUnprocessableEntity)
		return
	}

	group_id, err := sm.AddPricingGroup(name)
	log.Println("group created", group_id)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	//GET GROUP_ID DATA

	w.Write(SuccessResponse("Temporary success"))
}

func RenamePricingGroup(w http.ResponseWriter, r *http.Request) {
	storage, _ := app.GetStorage()

	id := chi.URLParam(r, "id")
	if id == "" {
		errMsg := "Id is required"
		http.Error(w, errMsg, http.StatusUnprocessableEntity)
		return
	}
	idInt, _ := strconv.Atoi(id)

	r.ParseForm()
	name := r.Form.Get("name")

	params := db.UpdatePricingGroupsParams{
		ID:   int32(idInt),
		Name: string(name),
	}

	err := storage.UpdatePricingGroups(context.Background(), params)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(SuccessResponse("Group successfully renamed"))
}

func DeletePricingGroup(w http.ResponseWriter, r *http.Request) {
	storage, _ := app.GetStorage()
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Id is required", http.StatusUnprocessableEntity)
		return
	}
	idInt, _ := strconv.Atoi(id)

	deleted_id, err := storage.DeletePricingGroup(context.Background(), int32(idInt))
	if deleted_id == 0 {
		http.Error(w, "No requested id in database", 500)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(SuccessResponse("Group successfully deleted"))
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
