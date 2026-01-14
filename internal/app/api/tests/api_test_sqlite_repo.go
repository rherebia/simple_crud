package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"learing_go/simple_crud/internal/app/api/handlers"
	"learing_go/simple_crud/internal/app/api/models"
	"learing_go/simple_crud/internal/app/api/repository"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetAlbumsWithSqlite(t *testing.T) {
	albumRepository := repository.NewAlbumJsonRepository()
	albumHandler := handlers.NewAlbumHandler(albumRepository)

	r := gin.Default()
	r.GET("/v1/albums", albumHandler.GetAlbums)

	expectedAlbums := []models.Album{
		{
			ID:     1,
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99,
		},
		{
			ID:     2,
			Title:  "Jeru",
			Artist: "Gerry Mulligan",
			Price:  17.99,
		},
		{
			ID:     3,
			Title:  "Sarah Vaughan and Clifford Brown",
			Artist: "Sarah Vaughan",
			Price:  39.99,
		},
	}

	req, _ := http.NewRequest(http.MethodGet, "/v1/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualAlbums []models.Album
	err := json.Unmarshal(w.Body.Bytes(), &actualAlbums)

	assert.NoError(t, err)
	assert.Len(t, actualAlbums, 3)
	assert.Equal(t, expectedAlbums, actualAlbums)
}

func TestShouldGetSingleAlbumWithSqlite(t *testing.T) {
	albumRepository := repository.NewAlbumJsonRepository()
	albumHandler := handlers.NewAlbumHandler(albumRepository)

	r := gin.Default()
	r.GET("/v1/albums/:id", albumHandler.GetAlbumByID)

	expectedAlbum := models.Album{
		ID:     1,
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	}

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/albums/%v", expectedAlbum.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualAlbum models.Album
	err := json.Unmarshal(w.Body.Bytes(), &actualAlbum)

	assert.NoError(t, err)
	assert.Equal(t, expectedAlbum, actualAlbum)
}

func TestShouldNotFindAlbumByIdWithSqlite(t *testing.T) {
	albumRepository := repository.NewAlbumJsonRepository()
	albumHandler := handlers.NewAlbumHandler(albumRepository)

	r := gin.Default()
	r.GET("/v1/albums/:id", albumHandler.GetAlbumByID)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/albums/%v", "4"), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var data map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &data)

	assert.NoError(t, err)
	assert.Equal(t, data["message"], "album not found")
}

func TestShouldCreateAlbumWithSqlite(t *testing.T) {
	albumRepository := repository.NewAlbumJsonRepository()
	albumHandler := handlers.NewAlbumHandler(albumRepository)

	r := gin.Default()
	r.GET("/v1/albums", albumHandler.GetAlbums)
	r.POST("/v1/albums", albumHandler.PostAlbums)

	newAlbum := models.Album{
		ID:     4,
		Title:  "On and On",
		Artist: "Jack Johnson",
		Price:  32.99,
	}

	initialAlbums := []models.Album{
		{
			ID:     1,
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99,
		},
		{
			ID:     2,
			Title:  "Jeru",
			Artist: "Gerry Mulligan",
			Price:  17.99,
		},
		{
			ID:     3,
			Title:  "Sarah Vaughan and Clifford Brown",
			Artist: "Sarah Vaughan",
			Price:  39.99,
		},
	}

	req, _ := http.NewRequest(http.MethodGet, "/v1/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualAlbums []models.Album
	err := json.Unmarshal(w.Body.Bytes(), &actualAlbums)

	assert.Equal(t, initialAlbums, actualAlbums)

	jsonData, _ := json.Marshal(newAlbum)
	req, _ = http.NewRequest(http.MethodPost, "/v1/albums", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var actualAlbum models.Album
	err = json.Unmarshal(w.Body.Bytes(), &actualAlbum)

	assert.NoError(t, err)
	assert.Equal(t, newAlbum, actualAlbum)

	finalAlbums := slices.Clone(initialAlbums)
	finalAlbums = append(finalAlbums, newAlbum)

	req, _ = http.NewRequest(http.MethodGet, "/v1/albums", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var albumsAfterCreation []models.Album
	err = json.Unmarshal(w.Body.Bytes(), &albumsAfterCreation)

	assert.Equal(t, finalAlbums, albumsAfterCreation)
}

func TestShouldUpdateAlbumWithSqlite(t *testing.T) {
	albumRepository := repository.NewAlbumJsonRepository()
	albumHandler := handlers.NewAlbumHandler(albumRepository)

	r := gin.Default()
	r.PUT("/v1/albums/:id", albumHandler.UpdateAlbum)

	updatedAlbum := models.Album{
		ID:     1,
		Title:  "Californication",
		Artist: "Red Hot Chili Peppers",
		Price:  40.99,
	}

	jsonData, _ := json.Marshal(updatedAlbum)
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/v1/albums/%v", "1"), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var data map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &data)

	assert.NoError(t, err)
	assert.Equal(t, data["message"], "album updated")
}

func TestShouldDeleteAlbumWithSqlite(t *testing.T) {
	albumRepository := repository.NewAlbumJsonRepository()
	albumHandler := handlers.NewAlbumHandler(albumRepository)

	r := gin.Default()
	r.DELETE("/v1/albums/:id", albumHandler.DeleteAlbum)

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/albums/%v", "1"), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var data map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &data)

	assert.NoError(t, err)
	assert.Equal(t, data["message"], "album deleted")
}

func TestShouldNotDeleteAbsentAlbumWithSqlite(t *testing.T) {
	albumRepository := repository.NewAlbumJsonRepository()
	albumHandler := handlers.NewAlbumHandler(albumRepository)

	r := gin.Default()
	r.DELETE("/v1/albums/:id", albumHandler.DeleteAlbum)

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/albums/%v", "4"), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var data map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &data)

	assert.NoError(t, err)
	assert.Equal(t, data["message"], "album not found")
}
