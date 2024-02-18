package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"space-management-system/app"
	"time"
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
	data := struct {
		Space_id   int32
		Time_to    int64
		Car_number string
		Fee        float32
	}{}
	err := decodeJSONBody(w, r, &data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println("data loaded", data)
	sm, _ := app.GetSpacesManager()
	err = sm.UpdateSpaceStatus(data.Space_id, "open", "pending")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	rm, _ := app.GetReservationManager()
	time := time.Unix(data.Time_to, 0).UTC()
	id, err := rm.AddNewReservation(data.Car_number, data.Space_id, data.Fee, time, "pending")
	if err != nil {
		http.Error(w, err.Error(), 500)
		sm.UpdateSpaceStatus(data.Space_id, "pending", "open")
		return
	}
	response := struct {
		Success bool  `json:"success"`
		ID      int32 `json:"id"`
	}{
		Success: true,
		ID:      id,
	}
	jsonData, _ := json.Marshal(response)
	w.Write(jsonData)
}

func ConfirmReservationPayment(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Space_id int32
	}{}
	err := decodeJSONBody(w, r, &data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	sm, _ := app.GetSpacesManager()
	err = sm.UpdateSpaceStatus(data.Space_id, "pending", "reserved")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	rm, _ := app.GetReservationManager()
	err = rm.ChangeReservationStatus(data.Space_id, "reserved")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(SuccessResponse("space was successfully reserved"))
}
