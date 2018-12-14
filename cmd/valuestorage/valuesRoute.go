package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/wutchzone/value-storage-service/pkg/value"
)

// HandleGet list all values
func HandleGet(w http.ResponseWriter, r *http.Request) {
	v := r.Context().Value(value.ValueKey).(value.Value)

	// read documents

	cursor, err := DB.Collection(v.Key).Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
	v := r.Context().Value(value.ValueKey).(value.Value)

	DB.Collection(v.Key).InsertOne(r.Context(), v)
}
