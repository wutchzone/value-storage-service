package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/wutchzone/value-storage-service/pkg/value"
)

// InitRoutes for value storage microservice
func InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", sayHello)
	r.Route("/api", func(r chi.Router) {
		r.Route("/{unit}", func(r chi.Router) {
			r.Use(ValueCtx)
			r.Use(FilterCtx)

			// Post new value to the :unit: list
			r.Post("/", HandlePost)
			// Get list of all value of type :unit:
			r.Get("/", HandleGet)

			r.Route("/{uuid}", func(r chi.Router) {
				// Manipulate with single record
				r.Get("/", HandleGetOne)
				r.Delete("/", HandleDeleteOne)
				// r.Put("/", HandleUpdateOne) Maybe it is not neccesarry to update values
			})
		})
	})

	return r
}

func sayHello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Server is active and listening on port: " + strconv.Itoa(Config.Port)))
}

// ValueCtx for parsing request data to app understable content
func ValueCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var v value.Value
		err := json.NewDecoder(r.Body).Decode(&v)
		v.Key = chi.URLParam(r, "unit")

		if v.Key == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		if err != nil && r.Method != "GET" && r.Method != "DELETE" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), value.ValueKey, v)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// FilterCtx for parsing request data to app understable content
func FilterCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			f := value.Filter{}

			if ffrom, err := time.Parse(time.RFC3339, strings.Split(strings.Replace(r.URL.Query().Get("from"), " ", "+", -1), "+")[0]); err == nil {
				f.From = &ffrom
			}
			if ftom, err := time.Parse(time.RFC3339, strings.Split(strings.Replace(r.URL.Query().Get("to"), " ", "+", -1), "+")[0]); err == nil {
				f.To = &ftom
			}

			ctx := context.WithValue(r.Context(), value.FilterKey, f)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
