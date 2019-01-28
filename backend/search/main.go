package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var searchService SearchService

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	searchService = SearchService{
		spotifyClient: &SpotifyImpl{},
	}

	searchService.AuthSpotify()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/search/{input}", GetSearch).Methods("GET")

	log.Println("Listening on :8083")
	log.Fatal(http.ListenAndServe(":8083", router))
}

func GetSearch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	items := searchService.SearchSong(params["input"])

	w.Header().Set("Access-Control-Allow-Origin", "*")
	_ = json.NewEncoder(w).Encode(items)

}
