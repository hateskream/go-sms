package handlers

import (
	"encoding/json"
	"net/http"
	"space-management-system/app"
	"strconv"
)

func SuccessResponse(msg string) []byte {
	response := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: msg,
	}
	result, _ := json.Marshal(response)
	return result
}

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
	sm, _ := app.GetSpacesManager()
	r.ParseForm()
	space_id, err := strconv.Atoi(r.Form.Get("space"))
	// car_number := r.Form.Get("car_number")

	// if car_number == "" {
	// 	http.Error(w, "car number is required parameter", 422)
	// }
	if err != nil {
		http.Error(w, "incorrect space_id", 422)
		return
	}
	err = sm.UpdateSpaceStatus(int32(space_id), "open", "pending")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(SuccessResponse("space was successfully reserved"))
}

func AddReservation(w http.ResponseWriter, r *http.Request) {
	rm, _ := app.GetReservationManager()
}

func ConfirmReservation(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	rm, _ := app.GetReservationManager()
	rm.ChangeReservationStatus(int32(id), "pending")
}

func UpdateReservationStatus(w http.ResponseWriter, r *http.Request) {
	sm, _ := app.GetSpacesManager()
	r.ParseForm()
	space_id, err := strconv.Atoi(r.Form.Get("space"))

	err = sm.UpdateSpaceStatus(int32(space_id), "pending", "reserved")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(SuccessResponse("space was successfully reserved"))
}
