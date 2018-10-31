package main

import "github.com/go-chi/chi"

// InitRoutes for value storage microservice
func InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/data", func(r chi.Router) {
		r.Route("/{unit}", func(r chi.Router) {
			// Post new value to the :unit: list
			r.Post("/", nil)
			// Get list of all value of type :unit:
			r.Get("/", nil)

			r.Route("/{uuid}", func(r chi.Router) {
				// Manipulate with single record
				r.Get("/", nil)
				r.Delete("/", nil)
				r.Put("/", nil)
			})
		})

	})

	return r
}
