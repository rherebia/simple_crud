package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbum)
	router.PUT("/albums/:id", updateAlbum)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(context *gin.Context) {
	id := context.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

// deleteAlbum deletes an album from the list.
func deleteAlbum(context *gin.Context) {
	id := context.Param("id")
	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			context.JSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// updateAlbum updates an album from JSON received in the request body.
func updateAlbum(context *gin.Context) {
	id := context.Param("id")
	var newAlbum album
	for i, album := range albums {
		if album.ID == id {
			if err := context.BindJSON(&newAlbum); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			albums[i] = newAlbum
			context.JSON(http.StatusOK, gin.H{"message": "album updated"})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
