package main

import (
	"learing_go/simple_crud/internal/app/api/handlers"
	"learing_go/simple_crud/internal/app/api/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	albumRepository := repository.NewAlbumJsonRepository()
	albumHandler := handlers.NewAlbumHandler(albumRepository)
	router.GET("/albums", albumHandler.GetAlbums)
	router.GET("/albums/:id", albumHandler.GetAlbumByID)
	router.POST("/albums", albumHandler.PostAlbums)
	router.DELETE("/albums/:id", albumHandler.DeleteAlbum)
	router.PUT("/albums/:id", albumHandler.UpdateAlbum)
	router.Run("localhost:8080")
}
