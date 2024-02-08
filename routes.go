package main

import (
	"space-management-system/app"

	"github.com/go-chi/chi/v5"
)

func initRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/admin/features", app.Handlers.GetAllSpaces)
	r.Route("/features", func(r chi.Router) {
		r.Get("/", app.Handlers.GetFeatures)
		r.Post("/add", app.Handlers.AddFeature) // PUT /features/123
		r.Route("/{featureID}", func(r chi.Router) {
			r.Put("/", app.Handlers.UpdateFeature)    // PUT /features/123
			r.Delete("/", app.Handlers.DeleteFeature) // DELETE /features/123
		})
	})
	return r
}
