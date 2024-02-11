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
			r.Put("/", app.Handlers.UpdateFeature)    // ВАСЯ!!
			r.Delete("/", app.Handlers.DeleteFeature) // РЕАЛИЗОВАНО
		})
	})
	r.Route("/spaces", func(r chi.Router) {
		r.Get("/", app.Handlers.GetSpaces)
		r.Post("/add", app.Handlers.AddSpace)
		r.Route("/{spaceID}", func(r chi.Router) {
			r.Put("/", app.Handlers.UpdateSpace)
			r.Delete("/", app.Handlers.DeleteSpace)
		})
	})

	return r
}
