package database

import (
	"MusicFinder/src/database/Album"
	"strconv"
	"testing"
)

func TestByArtist(t *testing.T) {
	albums, err := Album.ByArtist("Coltrane")
	if len(albums) == 0 {
		t.Errorf("ByArtist() empty set : awaited 2")
	}
	if err != nil {
		t.Errorf("ByArtist() error %v", err)
	}
}

func TestById(t *testing.T) {
	album, err := Album.ById(1)
	if album == nil {
		t.Errorf("ById() empty set.")
	}
	if err != nil {
		t.Errorf("ById() error %v", err)
	}
}

func TestCreate(t *testing.T) {
	createId, err := Album.Create(Album.Album{
		Title:  "TestTitle",
		Artist: "TestArtist",
		Price:  50.00,
	})
	if createId == 0 {
		t.Errorf("Create() returns 0 id.")
	}
	if err != nil {
		t.Errorf("Create error %v", err)
	}
	createdAlbum, err := Album.ById(createId)
	t.Logf("ID : %v, Title : %v, Album : %v, Price : %v", strconv.FormatInt(createdAlbum.ID, 10), createdAlbum.Title, createdAlbum.Artist, createdAlbum.Price)
}
