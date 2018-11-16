package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/wutchzone/value-storage-service/api"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	// DB Reference
	DB *mongo.Database
	// Config reference
	Config api.Config
)

func main() {
	// Load config
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Println("Failed to load config file.")
		os.Exit(1)
	}

	// Parse config
	f, err := ioutil.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("Can not read file", f)
		os.Exit(1)
	}
	json.NewDecoder(strings.NewReader(string(f))).Decode(&Config)

	// Init connection to the DB
	client, err := mongo.NewClient(Config.DbURL)
	if err != nil {
		fmt.Println("There was an error connecting to the DB ", err)
		os.Exit(1)
	}

	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("Client failed connect DB ", err)
		os.Exit(1)
	}

	DB = client.Database(Config.TableName)
	//DB.Collection("test").InsertOne(context.Background(), value.NewValue("test", "20"))

	// Init routes
	r := InitRoutes()

	fmt.Println("Server started.")
	http.ListenAndServe(":"+strconv.Itoa(Config.Port), r)
}
