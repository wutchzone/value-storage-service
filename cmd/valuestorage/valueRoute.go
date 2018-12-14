package main

import (
	"encoding/json"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/wutchzone/value-storage-service/pkg/value"

	"github.com/go-chi/chi"
)

// HandleGetOne gets record with specified UUID from URL
func HandleGetOne(w http.ResponseWriter, r *http.Request) {
	v := r.Context().Value(value.ValueKey).(value.Value)
	uid := chi.URLParam(r, "uuid")
	oid, _ := objectid.FromHex(uid)

	found := DB.Collection(v.Key).FindOne(r.Context(), bson.M{"_id": oid})
	found.Decode(&v)
	if v.ID.IsZero() {
		http.Error(w, "Document not found", http.StatusNotFound)
	} else {
		response, _ := json.Marshal(v)
		w.Write(response)
	}
}

// HandleUpdateOne updates corresponding record
func HandleUpdateOne(w http.ResponseWriter, r *http.Request) {
	v := r.Context().Value(value.ValueKey).(value.Value)
	uid := chi.URLParam(r, "uuid")

	oid, _ := objectid.FromHex(uid)
	if DB.Collection(v.Key).FindOneAndUpdate(r.Context(), bson.M{"_id": oid}, v) != nil {
		http.Error(w, "Document not found", http.StatusNotFound)
	}
}

// HandleDeleteOne updates corresponding record
func HandleDeleteOne(w http.ResponseWriter, r *http.Request) {
	v := r.Context().Value(value.ValueKey).(value.Value)
	uid := chi.URLParam(r, "uuid")

	oid, _ := objectid.FromHex(uid)
	if DB.Collection(v.Key).FindOneAndDelete(r.Context(), bson.M{"_id": oid}) != nil {
		http.Error(w, "Document not found", http.StatusNotFound)
	}
}
