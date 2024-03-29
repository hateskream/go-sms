package main

import (
	"space-management-system/handlers"

	"github.com/go-chi/chi/v5"
)

func initRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/spaces", func(r chi.Router) {
		r.Get("/", handlers.GetSpaces)
	})
	r.Route("/features", func(r chi.Router) {
		r.Get("/", handlers.GetFeatures) //РЕАЛИЗОВАНО
		// r.Post("/add", handlers.AddFeature) // РЕАЛИЗОВАНО
		// r.Route("/{featureID}", func(r chi.Router) {
		// 	r.Put("/", handlers.UpdateFeature)    // ВАСЯ!!
		// 	r.Delete("/", handlers.DeleteFeature) // РЕАЛИЗОВАНО
		// })
	})

	//reservations/history
	r.Route("/reservation", func(r chi.Router) {
		r.Post("/create", handlers.ReserveSpace)
		r.Post("/confirm", handlers.ConfirmReservationPayment)
	})

	r.Route("/pricing_groups", func(r chi.Router) {
		r.Post("/add", handlers.AddPricingGroup)
		// r.Route("/{pricing_groupID}", func(r chi.Router) {
		// 	r.Put("/", handlers.RenamePricingGroup)
		// 	r.Delete("/", handlers.DeletePricingGroup)
		// })
	})

	r.Route("/admin", func(r chi.Router) {
		r.Route("/spaces", func(r chi.Router) {
			r.Post("/add", handlers.AddSpace)
			r.Post("/features", handlers.UpdateSpaceFeatures)
		})
		r.Route("/reservations", func(r chi.Router) {
			r.Get("/", handlers.GetReservationHistory)
			r.Get("/active", handlers.GetActiveReservations)
		})
	})

	r.Route("/parking", func(r chi.Router) {
		r.Patch("/reservations/{reservationID}", handlers.CarDeparture)
	})

	return r
}
