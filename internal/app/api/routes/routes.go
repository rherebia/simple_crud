package routes

import (
	"learing_go/simple_crud/internal/app/api/handlers"
	"learing_go/simple_crud/internal/app/api/repository"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	v1 := server.Group("/v1")

	albumRepository := repository.NewAlbumJsonRepository()
	albumHandler := handlers.NewAlbumHandler(albumRepository)

	v1.GET("/albums", albumHandler.GetAlbums)
	v1.GET("/albums/:id", albumHandler.GetAlbumByID)
	v1.POST("/albums", albumHandler.PostAlbums)
	v1.DELETE("/albums/:id", albumHandler.DeleteAlbum)
	v1.PUT("/albums/:id", albumHandler.UpdateAlbum)
}
