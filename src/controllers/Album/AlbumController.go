package Album

import (
	"MusicFinder/src/database/Album"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAlbums(c *gin.Context) {
	albums, err := Album.All()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "An error occurred.")
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum Album.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "An error occurred.")
		panic(err)
	}
	createId, err := Album.Create(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "An error occurred.")
		panic(err)
	}
	newAlbum.ID = createId
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetOneById(c *gin.Context) {
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "An error occurred.")
		panic(err)
	}
	album, err := Album.ById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "No album found for this ID.")
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, album)
}
