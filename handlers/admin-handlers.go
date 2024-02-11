package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"space-management-system/services/db/db"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handlers) GetSpacesInfo(w http.ResponseWriter, r *http.Request) {
	//get space id, name, status, group and list of features + pagination
}

func (h *Handlers) AddSpace(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	if name == "" {
		http.Error(w, "name is required field", 422)
		return
	}
	physical_id := r.Form.Get("physical_id")
	physical_idInt, _ := strconv.Atoi(physical_id)
	group_id := r.Form.Get("group_id")
	group_idInt, _ := strconv.Atoi(group_id)
	status_id := r.Form.Get("status_id")
	status_idInt, _ := strconv.Atoi(status_id)

	params := db.AddSpaceParams{
		Name:       name,
		PhysicalID: pgtype.Int4{Int32: int32(physical_idInt), Valid: true},
		GroupID:    pgtype.Int4{Int32: int32(group_idInt), Valid: true},
		StatusID:   pgtype.Int4{Int32: int32(status_idInt), Valid: true},
	}

	log.Printf("GetRequest: %s %s %s %s", name, physical_id, group_id, status_id)
	added, err := h.Storage.AddSpace(context.Background(), params)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	jsonData, err := json.Marshal(added)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write(jsonData)
}

func (h *Handlers) DeleteSpace(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "spaceID")
	if id == "" {
		errMsg := "Id is required"
		http.Error(w, errMsg, http.StatusUnprocessableEntity)
		return
	}
	idInt, _ := strconv.Atoi(id)
	deleted_id, err := h.Storage.DeleteSpace(context.Background(), int32(idInt))
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

func (h *Handlers) UpdateSpace(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "spaceID")
	if id == "" {
		errMsg := "Id is required"
		http.Error(w, errMsg, http.StatusUnprocessableEntity)
		return
	}
	idInt, _ := strconv.Atoi(id)
	r.ParseForm()
	name := r.Form.Get("name")
	log.Printf("id: %s %s", id, name)
	if name == "" {
		http.Error(w, "name is required field", 422)
		return
	}
	physical_id := r.Form.Get("physical_id")
	physical_idInt, _ := strconv.Atoi(physical_id)
	group_id := r.Form.Get("group_id")
	group_idInt, _ := strconv.Atoi(group_id)
	status_id := r.Form.Get("status_id")
	status_idInt, _ := strconv.Atoi(status_id)

	params := db.UpdateSpaceParams{
		ID:         int32(idInt),
		Name:       name,
		PhysicalID: pgtype.Int4{Int32: int32(physical_idInt), Valid: true},
		GroupID:    pgtype.Int4{Int32: int32(group_idInt), Valid: true},
		StatusID:   pgtype.Int4{Int32: int32(status_idInt), Valid: true},
	}

	err := h.Storage.UpdateSpace(context.Background(), params)

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
		Message: "Feature successfully updated.",
	}
	jsonData, _ := json.Marshal(response)
	w.Write(jsonData)
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

	if id == "" {
		errMsg := "Id is required"
		http.Error(w, errMsg, http.StatusUnprocessableEntity)
		return
	}
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

	id := chi.URLParam(r, "featureID")

	if id == "" {
		errMsg := "Id is required"
		http.Error(w, errMsg, http.StatusUnprocessableEntity)
		return
	}

	idInt, _ := strconv.Atoi(id)

	r.ParseForm()
	name := r.Form.Get("name")

	params := db.UpdateFeatureParams{
		ID:   int32(idInt),
		Name: name,
	}
	err := h.Storage.UpdateFeature(context.Background(), params)

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
		Message: "Feature successfully updated.",
	}
	jsonData, _ := json.Marshal(response)
	w.Write(jsonData)

}

func (h *Handlers) UpdatePricingPolicy(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "featureID")
	// idInt, _ := strconv.Atoi(id)

	// r.ParseForm()
	// name := r.Form.Get("name")
}

func (h *Handlers) AddPricingGroup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	group_id, err := h.Storage.AddPricingGroup(context.Background(), name)

	if err != nil {
		errMsg := err.Error()
		http.Error(w, errMsg, 500)
		return
	}
	rate := 1

	params := db.GeneratTimePricingPolicyParams{
		GroupID: int32(group_id),
		Rate:    float32(rate),
	}

	errGenerate := h.Storage.GeneratTimePricingPolicy(context.Background(), params)
	if errGenerate != nil {
		errGenerateMsg := errGenerate.Error()
		http.Error(w, errGenerateMsg, 500)
		return
	}

	response := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "Group successfully added.",
	}
	jsonData, _ := json.Marshal(response)
	w.Write(jsonData)

}

func (h *Handlers) DeletePricingGroup(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "pricing_groupID")
	if id == "" {
		errMsg := "Id is required"
		http.Error(w, errMsg, http.StatusUnprocessableEntity)
		return
	}
	idInt, _ := strconv.Atoi(id)

	errClean := h.Storage.CleanTimePricingPolicy(context.Background(), int32(idInt))
	if errClean != nil {
		errMsg := errClean.Error()
		http.Error(w, errMsg, 500)
		return
	}

	group_id, err := h.Storage.DeletePricingGroup(context.Background(), int32(idInt))
	if err != nil {
		errMsg := err.Error()
		http.Error(w, errMsg, 500)
		return
	}
	log.Printf("Deleted: %d", group_id)
	response := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "Group successfully deleted.",
	}
	jsonData, _ := json.Marshal(response)
	w.Write(jsonData)
}

func (h *Handlers) UpdatePricingGroups(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "pricing_groupID")
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
	log.Println(params)
	err := h.Storage.UpdatePricingGroups(context.Background(), params)
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
		Message: "Group successfully updated.",
	}
	jsonData, _ := json.Marshal(response)
	w.Write(jsonData)

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
