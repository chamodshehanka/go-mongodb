package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shehanka/go-mongodb/models"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

//type Podcast struct {
//	ID     primitive.ObjectID `bson:"_id,omitempty"`
//	Name   string             `bson:"name,omitempty"`
//	Author string             `bson:"author,omitempty"`
//	Tags   []string           `bson:"tags,omitempty"`
//}

// Init podcasts
var podcasts []models.Podcast

func getPodcasts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(podcasts)
}

func loadPodcastsList()  {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	user := viper.GetString("database.user")
	password := viper.GetString("database.password")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + user + ":" + password + "@cluster0.rkyph.mongodb.net/cluster0?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatalf("%s", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("%s", err)
	}

	testDB := client.Database("cluster0")
	podcastsCollection := testDB.Collection("podcasts")

	cursor,err := podcastsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("%s", err)
	}

	var pods []models.Podcast
	if err = cursor.All(ctx, &pods); err != nil {
		log.Fatal(err)
	}
	podcasts = append(podcasts, pods[1])
}

func main() {
	loadPodcastsList()
	// Init router
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/podcast", handler).Methods("POST")
	r.HandleFunc("/podcast",handler).Methods("PUT")
	r.HandleFunc("/podcast/list",getPodcasts).Methods("GET")
	r.HandleFunc("/episodes", handler).Methods("GET")
	r.HandleFunc("/articles/{id}", handler).Methods("GET", "PUT")
}
