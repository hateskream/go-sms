package reservations

import (
	"context"
	"errors"
	"fmt"
	"log"
	"space-management-system/app"
	"space-management-system/services/db/db"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type ReservationsManager struct {
	App                 *app.App
	onUpdate            func(*app.ActiveReservation)
	active_reservations []app.ActiveReservation
	statuses            []db.ReservationStatus
}

func (rm *ReservationsManager) SetApp(app *app.App) error {
	rm.App = app
	rm.GetActiveReservations()
	return nil
}

func (rm *ReservationsManager) reservationUpdated(reservation *app.ActiveReservation) {
	log.Println("reservation updated", reservation)
	if rm.onUpdate != nil {
		rm.onUpdate(reservation)
	}
}

func (rm *ReservationsManager) getStatusID(name string) int32 {
	for _, status := range rm.statuses {
		if status.Name == name {
			return status.ID
		}
	}
	return -1
}

func (rm *ReservationsManager) getStatusName(id int32) string {
	for _, status := range rm.statuses {
		if status.ID == id {
			return status.Name
		}
	}
	return ""
}

func (rm *ReservationsManager) convertReservationRow(row db.GetActiveReservationsRow) app.ActiveReservation {
	reservation := app.ActiveReservation{
		ID: row.ID, TimeFrom: row.TimeFrom.Time, TimeTo: row.TimeTo.Time,
		CarID: row.CarID.Int32, SpaceID: row.SpaceID.Int32,
		Status: rm.getStatusName(row.StatusID.Int32),
	}
	return reservation
}

func (rm *ReservationsManager) GetCarId(car_number string) (int32, error) {
	log.Println("car find", car_number)
	id, _ := rm.App.Storage.GetCardByNumber(context.Background(), car_number)
	if id != 0 {
		return id, nil
	}
	id, err := rm.App.Storage.AddCardNumber(context.Background(), car_number)
	log.Println("car_id_insert", id, err)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (rm *ReservationsManager) GetActiveReservations() ([]app.ActiveReservation, error) {
	if len(rm.active_reservations) > 0 {
		return rm.active_reservations, nil
	}

	reservations, err := rm.App.Storage.GetActiveReservations(context.Background())
	if err != nil {
		return nil, err
	}
	rm.UpdateReservationStatuses()
	for _, r := range reservations {
		rm.active_reservations = append(rm.active_reservations, rm.convertReservationRow(r))
	}
	log.Println("reservations loaded", rm.active_reservations)
	return rm.active_reservations, nil
}

func (rm *ReservationsManager) AddNewReservation(car_number string, space_id int32, rservation_fee float32, time_to time.Time, status string) (int32, error) {
	rm.UpdateReservationStatuses()
	log.Println("statuses updated", rm.statuses)
	status_id := rm.getStatusID(status)
	if status_id == -1 {
		return 0, fmt.Errorf("No %s status", status)
	}
	log.Println("status", status_id)

	car_id, err := rm.GetCarId(car_number)
	if err != nil {
		return 0, err
	}
	currentTimeUTC := time.Now().UTC()
	log.Println("car", car_id)
	params := db.AddReservationParams{
		StatusID:       pgtype.Int4{Int32: status_id, Valid: true},
		SpaceID:        pgtype.Int4{Int32: space_id, Valid: true},
		ReservationFee: rservation_fee,
		TimeFrom:       pgtype.Timestamp{Time: currentTimeUTC, Valid: true},
		TimeTo:         pgtype.Timestamp{Time: time_to, Valid: true},
		CarID:          pgtype.Int4{Int32: car_id, Valid: true},
	}
	log.Println("params", params)
	id, err := rm.App.Storage.AddReservation(context.Background(), params)
	log.Println("result", id, err)
	if err != nil {
		return 0, err
	}
	new_reservation := app.ActiveReservation{
		ID: id, TimeFrom: currentTimeUTC, TimeTo: time_to,
		CarID: car_id, SpaceID: space_id, Status: status,
	}
	rm.active_reservations = append(rm.active_reservations, new_reservation)
	rm.reservationUpdated(&new_reservation)
	return id, nil
}

func (rm *ReservationsManager) ChangeReservationStatus(id int32, status string) error {
	rm.UpdateReservationStatuses()
	status_id := rm.getStatusID(status)

	if status_id == -1 {
		return errors.New("status not found")
	}

	params := db.UpdateReservationStatusParams{
		ID:       id,
		StatusID: pgtype.Int4{Int32: status_id, Valid: true},
	}

	err := rm.App.Storage.UpdateReservationStatus(context.Background(), params)

	for i, res := range rm.active_reservations {
		if res.ID == id {
			rm.active_reservations[i].Status = status
			rm.reservationUpdated(&rm.active_reservations[i])
			break
		}
	}
	return err
}

func (rm *ReservationsManager) UpdateReservationStatuses() error {
	if len(rm.statuses) == 0 {
		statuses, err := rm.App.Storage.GetReservationStatuses(context.Background())
		if err != nil {
			return err
		}
		rm.statuses = statuses
	}
	return nil
}

func (rm *ReservationsManager) GetReservationHistory(time_from time.Time, time_to time.Time, page int32, per_page int32) ([]db.GetReservationsHistoryRow, int32, error) {
	timeFilter := db.GetReservationsHistoryCountParams{
		TimeFrom:   pgtype.Timestamp{Time: time_from, Valid: true},
		TimeFrom_2: pgtype.Timestamp{Time: time_to, Valid: true},
	}
	count, err := rm.App.Storage.GetReservationsHistoryCount(context.Background(), timeFilter)
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, nil
	}

	params := db.GetReservationsHistoryParams{
		Limit:      per_page,
		Offset:     (page - 1) * per_page,
		TimeFrom:   timeFilter.TimeFrom,
		TimeFrom_2: timeFilter.TimeFrom_2,
	}

	reservations, err := rm.App.Storage.GetReservationsHistory(context.Background(), params)
	if err != nil {
		return nil, 0, err
	}
	return reservations, int32(count), nil

}

func (rm *ReservationsManager) UpdateReservationParkingData(data db.UpdateReservationParkingDataParams) error {
	log.Println(data)
	err := rm.App.Storage.UpdateReservationParkingData(context.Background(), data)
	return err
}
