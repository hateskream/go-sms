package reservations

import (
	"context"
	"errors"
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

func (rm *ReservationsManager) GetCarId(car_number string) (int32, error) {
	id, err := rm.App.Storage.GetCardByNumber(context.Background(), car_number)
	if err != nil {
		return 0, err
	}
	if id != 0 {
		return id, nil
	}
	id, err = rm.App.Storage.AddCardNumber(context.Background(), car_number)
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
	statusesMap := make(map[int32]string)
	for _, s := range rm.statuses {
		statusesMap[s.ID] = s.Name
	}
	for _, r := range reservations {
		rm.active_reservations = append(rm.active_reservations, app.ActiveReservation{
			GetActiveReservationsRow: r,
			Status:                   statusesMap[r.StatusID.Int32],
		})
	}
	return rm.active_reservations, nil
}

func (rm *ReservationsManager) AddNewReservation(car_number string, space_id int32, rservation_fee float32, time_to time.Time) (int32, error) {
	rm.UpdateReservationStatuses()
	status_id := int32(0)
	for _, status := range rm.statuses {
		if status.Name == "pending" {
			status_id = int32(status.ID)
			break
		}
	}
	if status_id == 0 {
		return 0, errors.New("No pending status")
	}

	car_id, err := rm.GetCarId(car_number)
	if err != nil {
		return 0, err
	}
	currentTimeUTC := time.Now().UTC()

	params := db.AddReservationParams{
		StatusID:       pgtype.Int4{Int32: status_id, Valid: true},
		SpaceID:        pgtype.Int4{Int32: space_id, Valid: true},
		ReservationFee: rservation_fee,
		TimeFrom:       pgtype.Timestamp{Time: currentTimeUTC, Valid: true},
		TimeTo:         pgtype.Timestamp{Time: time_to, Valid: true},
		CarID:          pgtype.Int4{Int32: car_id, Valid: true},
	}

	id, err := rm.App.Storage.AddReservation(context.Background(), params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (rm *ReservationsManager) ChangeReservationStatus(id int32, status string) error {
	rm.UpdateReservationStatuses()
	status_id := int32(0)
	for _, status := range rm.statuses {
		if status.Name == "pending" {
			status_id = int32(status.ID)
			break
		}
	}
	if status_id == 0 {
		return errors.New("Status not found")
	}

	params := db.UpdateReservationStatusParams{
		ID:       id,
		StatusID: pgtype.Int4{Int32: status_id, Valid: true},
	}

	err := rm.App.Storage.UpdateReservationStatus(context.Background(), params)
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

func (rm *ReservationsManager) GetReservationHistory() ([]string, error) {
	return nil, nil
}
