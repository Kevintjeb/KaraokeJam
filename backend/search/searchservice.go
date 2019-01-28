package main

import "github.com/zmb3/spotify"

// The track type
type Track struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Artist     []Artist `json:"artist,omitempty"`
	Album      Album    `json:"album,omitempty"`
	Popularity int      `json:"popularity,omitempty"`
	Duration   int      `json:"duration,omitempty"`
	Year       string   `json:"year,omitempty"`
}
type Artist struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Image Image  `json:"image,omitempty"`
}
type Album struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Artist []Artist `json:"artist,omitempty"`
	Cover  Image    `json:"cover,omitempty"`
}
type Image struct {
	Small  string `json:"64,omitempty"`
	Medium string `json:"300,omitempty"`
	Large  string `json:"640,omitempty"`
}

type SearchService struct {
	spotifyClient SpotifyClient
}

func (service *SearchService) AuthSpotify() error {
	return service.spotifyClient.Authenticate()
}

func (service *SearchService) SearchSong(input string) []Track {
	return service.spotifyClient.Search(input)
}

func parseTracks(tracks *spotify.FullTrackPage) []Track {
	var items []Track

	// handle playlist results
	if tracks != nil {

		for _, item := range tracks.Tracks {
			var track Track

			track.ID = item.ID.String()
			track.Name = item.Name
			track.Popularity = item.Popularity
			track.Duration = item.Duration
			track.Year = item.Album.ReleaseDate

			var artist []Artist
			for _, artistItem := range item.Artists {
				artist = append(artist, Artist{artistItem.ID.String(), artistItem.Name, Image{}})
			}
			track.Artist = artist

			var album Album
			album.ID = item.Album.ID.String()
			album.Artist = artist
			album.Name = item.Album.Name

			var cover Image
			cover.Small = item.Album.Images[2].URL
			cover.Medium = item.Album.Images[1].URL
			cover.Large = item.Album.Images[0].URL

			album.Cover = cover

			track.Album = album

			items = append(items, Track{track.ID, track.Name, track.Artist, track.Album, track.Popularity, track.Duration, track.Year})

		}

	}

	return items
}
