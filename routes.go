package main

import (
	"space-management-system/handlers"

	"github.com/go-chi/chi/v5"
)

func initRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/spaces", handlers.GetSpaces)
	r.Route("/features", func(r chi.Router) {
		r.Get("/", handlers.GetFeatures) //РЕАЛИЗОВАНО
		// r.Post("/add", handlers.AddFeature) // РЕАЛИЗОВАНО
		// r.Route("/{featureID}", func(r chi.Router) {
		// 	r.Put("/", handlers.UpdateFeature)    // ВАСЯ!!
		// 	r.Delete("/", handlers.DeleteFeature) // РЕАЛИЗОВАНО
		// })
	})
	r.Route("/reservation", func(r chi.Router) {
		r.Post("/add", handlers.ReserveSpace)
	})

	return r
}
