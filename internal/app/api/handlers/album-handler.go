package handlers

import (
	"learing_go/simple_crud/internal/app/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
}

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (h *AlbumHandler) GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (h *AlbumHandler) GetAlbumByID(context *gin.Context) {
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
func (h *AlbumHandler) PostAlbums(context *gin.Context) {
	var newAlbum models.Album

	if err := context.BindJSON(&newAlbum); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

// deleteAlbum deletes an album from the list.
func (h *AlbumHandler) DeleteAlbum(context *gin.Context) {
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
func (h *AlbumHandler) UpdateAlbum(context *gin.Context) {
	id := context.Param("id")
	var newAlbum models.Album
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
