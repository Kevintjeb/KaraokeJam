package main

type SpotifyClient interface {
	Authenticate() error
	Search(input string) []Track
}
