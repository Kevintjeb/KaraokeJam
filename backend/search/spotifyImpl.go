package main

import (
	"context"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type SpotifyImpl struct {
	Spotify spotify.Client
}

func (client *SpotifyImpl) Authenticate() error {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
		return err
	}
	client.Spotify = spotify.Authenticator{}.NewClient(token)

	return nil
}

func (client *SpotifyImpl) Search(input string) []Track {
	results, err := client.Spotify.Search(input, spotify.SearchTypeTrack)

	if err != nil {
		err = client.Authenticate()
		if err != nil {
			log.Fatal(err)
			return nil
		}
		results, err = client.Spotify.Search(input, spotify.SearchTypeTrack)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}

	return parseTracks(results.Tracks)
}
