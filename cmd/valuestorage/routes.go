package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/wutchzone/value-storage-service/pkg/value"
)

// InitRoutes for value storage microservice
func InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", sayHello)
	r.Route("/api/data", func(r chi.Router) {
		r.Route("/{unit}", func(r chi.Router) {
			r.Use(ValueCtx)
			r.Use(FilterCtx)

			// Post new value to the :unit: list
			r.Post("/", HandlePost)
			// Get list of all value of type :unit:
			r.Get("/", HandleGet)

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

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is active and listening on port: " + strconv.Itoa(Config.Port)))
}

// ValueCtx for parsing request data to app understable content
func ValueCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var v value.Value
		err := json.NewDecoder(r.Body).Decode(&v)

		v.Key = chi.URLParam(r, "unit")
		//fmt.Println("Here", v)
		if v.Key == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		if err != nil && r.Method != "GET" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//fmt.Println("Here", v)
		ctx := context.WithValue(r.Context(), value.ValueKey, v)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// FilterCtx for parsing request data to app understable content
func FilterCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var v value.Value
		// err := json.NewDecoder(r.Body).Decode(&v)

		// if err != nil {
		// 	http.Error(w, http.StatusText(http.StatusBadRequest), 404)
		// 	return
		// }

		// ctx := context.WithValue(r.Context(), value.ValueKey, v)
		next.ServeHTTP(w, r)
	})
}
