package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/wutchzone/value-storage-service/pkg/value"
)

// HandleGet list all values
func HandleGet(w http.ResponseWriter, r *http.Request) {

	// read documents
	cursor, err := DB.Collection("inventory").Find(r.Context(), bson.NewDocument())
	if err != nil {
		log.Fatal(err)
	}

	itemsToSend := []value.Value{}
	itemRead := value.Value{}
	for cursor.Next(r.Context()) {
		err := cursor.Decode(&itemRead)
		if err != nil {
			log.Fatal(err)
		}
		itemsToSend = append(itemsToSend, itemRead)
	}
	response, _ := json.Marshal(itemsToSend)
	w.Write(response)
}

// HandlePost saves new dato to the DB
func HandlePost(w http.ResponseWriter, r *http.Request) {
	v := chi.URLParam(r, "unit")

	if !(len(v) > 0) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
