package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func getPodcasts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("")
}

func main() {
	// Init router
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/podcast", handler).Methods("POST")
	r.HandleFunc("/podcast",handler).Methods("PUT")
	r.HandleFunc("/podcast/list",handler).Methods("GET")
	r.HandleFunc("/episodes", handler).Methods("GET")
	r.HandleFunc("/articles/{id}", handler).Methods("GET", "PUT")
}
