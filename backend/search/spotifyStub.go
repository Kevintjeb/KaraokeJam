package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/zmb3/spotify"
)

type SpotifyStub struct {
}

func (client SpotifyStub) Authenticate() error {
	return nil
}

func (client SpotifyStub) Search(input string) []Track {
	var items spotify.FullTrackPage
	response, _ := ioutil.ReadFile("./spotify-response.json")
	err := json.Unmarshal(response, &items)

	if err != nil {
		panic(err)
	}

	return parseTracks(&items)
}
