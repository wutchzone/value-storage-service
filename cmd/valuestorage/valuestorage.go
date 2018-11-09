package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	DB *mongo.Database
)

func main() {
	r := InitRoutes()

	client, _ := mongo.NewClient("mongodb://localhost:27017")
	client.Connect(context.Background())
	DB = client.Database("mongosample")

	fmt.Println("Server started")

	http.ListenAndServe(":7081", r)
}
