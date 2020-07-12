package main

import (
	"encoding/json"
	"github.com/Shehanka/go-mongodb/models"
	"github.com/gorilla/mux"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

// Init podcasts
var podcasts []models.Podcast

func getPodcasts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(podcasts)
}

func loadPodcastsList()  {


	//client :=
	//testDB := client.Database("cluster0")
	//podcastsCollection := testDB.Collection("podcasts")
	//
	//cursor,err := podcastsCollection.Find(ctx, bson.M{})
	//if err != nil {
	//	log.Fatalf("%s", err)
	//}
	//
	//var pods []models.Podcast
	//if err = cursor.All(ctx, &pods); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(pods)
	//podcasts = append(podcasts, pods[1])
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
