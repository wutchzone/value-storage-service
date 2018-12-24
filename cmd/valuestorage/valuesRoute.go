package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/wutchzone/value-storage-service/pkg/value"
)

// HandleGet list all values
func HandleGet(w http.ResponseWriter, r *http.Request) {
	v := r.Context().Value(value.ValueKey).(value.Value)
	f := r.Context().Value(value.FilterKey).(value.Filter)
	if f.From == nil {
		f.From = &time.Time{}
	}
	if f.To == nil {
		ct := time.Now()
		f.To = &ct
	}

	fmt.Println("From ", f.From)
	fmt.Println("To: ", f.To)
	// read documents
	cursor, err := DB.Collection(v.Key).Find(r.Context(), bson.M{"created_at": bson.M{"$gt": f.From, "$lt": f.To}})
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
	v.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	v.ModifiedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	DB.Collection(v.Key).InsertOne(r.Context(), v)
}
