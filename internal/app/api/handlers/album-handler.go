package handlers

import (
	"learing_go/simple_crud/internal/app/api/models"
	"learing_go/simple_crud/internal/app/api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	repo repository.AlbumRepository
}

func NewAlbumHandler(r repository.AlbumRepository) AlbumHandler {
	return AlbumHandler{repo: r}
}

func (h *AlbumHandler) GetAlbums(context *gin.Context) {
	albums, err := h.repo.GetAlbums()

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Album not found",
		})
		return
	}

	context.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (h *AlbumHandler) GetAlbumByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not parse album id"})
		return
	}

	album, err := h.repo.GetAlbumById(id)

	if err != nil {
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

	err := h.repo.CreateAlbum(&newAlbum)

	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Could not create album. Try again later"})
		return
	}

	context.IndentedJSON(http.StatusCreated, newAlbum)
}

// deleteAlbum deletes an album from the list.
func (h *AlbumHandler) DeleteAlbum(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not parse album id"})
		return
	}

	err = h.repo.DeleteAlbum(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "album deleted"})
}

// updateAlbum updates an album from JSON received in the request body.
func (h *AlbumHandler) UpdateAlbum(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not parse album id"})
		return
	}

	var newAlbum models.Album
	if err := context.BindJSON(&newAlbum); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAlbum.ID = id

	err = h.repo.UpdateAlbum(newAlbum)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "album updated"})
}
