package handlers

import (
	"learing_go/simple_crud/internal/app/api/models"
	"learing_go/simple_crud/internal/app/api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	repo repository.AlbumRepository
}

func NewAlbumHandler(r repository.AlbumRepository) AlbumHandler {
	return AlbumHandler{repo: r}
}

func (h *AlbumHandler) GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, h.repo.GetAlbums())
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (h *AlbumHandler) GetAlbumByID(context *gin.Context) {
	id := context.Param("id")

	album := h.repo.GetAlbumById(id)

	if album.ID == "" {
		context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, album)
}

// postAlbums adds an album from JSON received in the request body.
func (h *AlbumHandler) PostAlbums(context *gin.Context) {
	var newAlbum models.Album

	if err := context.BindJSON(&newAlbum); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.repo.CreateAlbum(newAlbum)

	context.IndentedJSON(http.StatusCreated, newAlbum)
}

// deleteAlbum deletes an album from the list.
func (h *AlbumHandler) DeleteAlbum(context *gin.Context) {
	id := context.Param("id")

	deleted := h.repo.DeleteAlbum(id)

	if deleted {
		context.JSON(http.StatusOK, gin.H{"message": "album deleted"})
		return
	}

	context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// updateAlbum updates an album from JSON received in the request body.
func (h *AlbumHandler) UpdateAlbum(context *gin.Context) {
	id := context.Param("id")

	var newAlbum models.Album
	if err := context.BindJSON(&newAlbum); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAlbum.ID = id

	updated := h.repo.UpdateAlbum(newAlbum)

	if updated {
		context.JSON(http.StatusOK, gin.H{"message": "album updated"})
		return
	}

	context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
