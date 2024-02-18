package handlers

import (
	"net/http"
	"space-management-system/app"
	"space-management-system/services/db/db"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func CarArrival(w http.ResponseWriter, r *http.Request) {
	//change space status?
}

func CarDeparture(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "reservationID")
	data := struct {
		Time_from     int64
		Time_to       int64
		Fee           float32
		Breakdown_fee float32
	}{}
	err := decodeJSONBody(w, r, &data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	idInt, _ := strconv.Atoi(id)
	rm, _ := app.GetReservationManager()
	time_from := time.Unix(data.Time_from, 0).UTC()
	time_to := time.Unix(data.Time_to, 0).UTC()

	params := db.UpdateReservationParkingDataParams{
		ID:                  int32(idInt),
		ParkingTimeFrom:     pgtype.Timestamp{Time: time_from, Valid: true},
		ParkingTimeTo:       pgtype.Timestamp{Time: time_to, Valid: true},
		ParkingFee:          data.Fee,
		ParkingFeeBreakdown: data.Breakdown_fee,
	}
	err = rm.UpdateReservationParkingData(params)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(SuccessResponse("Reservation parking data was successfully updated"))
}
