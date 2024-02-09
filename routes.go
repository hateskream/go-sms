package main

import (
	"space-management-system/app"

	"github.com/go-chi/chi/v5"
)

func initRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/places", app.Handlers.GetSpaces)
	r.Route("/features", func(r chi.Router) {
		r.Get("/", app.Handlers.GetFeatures)    //РЕАЛИЗОВАНО
		r.Post("/add", app.Handlers.AddFeature) // РЕАЛИЗОВАНО
		r.Route("/{featureID}", func(r chi.Router) {
			r.Patch("/", app.Handlers.UpdateFeature)  // ВАСЯ!!
			r.Delete("/", app.Handlers.DeleteFeature) // РЕАЛИЗОВАНО
		})
	})
	r.Route("/spaces", func(r chi.Router) {
		r.Get("/", app.Handlers.GetSpaces)    //РЕАЛИЗОВАНО
		r.Post("/add", app.Handlers.AddSpace) // РЕАЛИЗОВАНО
		r.Route("/{placeID}", func(r chi.Router) {
			r.Patch("/", app.Handlers.UpdateSpace)  // ВАСЯ!!
			r.Delete("/", app.Handlers.DeleteSpace) // РЕАЛИЗОВАНО
		})
	})

	return r
}
