package main

import (
	"MusicFinder/src/controllers/Album"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", Album.GetAlbums)
	router.GET("/album/:id", Album.GetOneById)
	router.POST("/album", Album.PostAlbum)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
