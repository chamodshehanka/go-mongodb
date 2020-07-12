package main

import (
	"context"
	"encoding/json"
	"github.com/Shehanka/go-mongodb/config"
	"github.com/Shehanka/go-mongodb/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

// Init podcasts, episodes
var podcasts []models.Podcast
var episodes []models.Episode

func getPodcasts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(podcasts)
}

func getEpisodes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episodes)
}

func loadPodcastsList() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := config.GetMongoDBConnection()

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
	podcasts = pods
}

func main() {
	loadPodcastsList()
	// Init router
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/podcast", handler).Methods("POST")
	r.HandleFunc("/podcast", handler).Methods("PUT")
	r.HandleFunc("/podcast/list", getPodcasts).Methods("GET")
	r.HandleFunc("/episodes/list", getEpisodes).Methods("GET")
	r.HandleFunc("/articles/{id}", handler).Methods("GET", "PUT")

	log.Println(http.ListenAndServe(":7000", r))
}
