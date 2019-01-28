package main

import "testing"

func TestSearchSpotify(t *testing.T) {
	service := SearchService{
		spotifyClient: &SpotifyStub{},
	}

	t.Run("Auth spotify, search song", func(t *testing.T) {

		err := service.AuthSpotify()

		if err != nil {
			t.FailNow()
		}

		tracks := service.SearchSong("trackname")

		if tracks == nil {
			t.FailNow()
		}

		if tracks[0].Artist[0].Name == "" {
			t.FailNow()
		}

		if tracks[0].Name == "" {
			t.FailNow()
		}

		if tracks[0].Album.Name == "" {
			t.FailNow()
		}

		if tracks[0].Album.Cover.Large == "" {
			t.FailNow()
		}

		if tracks[0].Duration <= 0 {
			t.FailNow()
		}

		if tracks[0].Year == "" {
			t.FailNow()
		}

		if err != nil {
			t.FailNow()
		}

	})
}
