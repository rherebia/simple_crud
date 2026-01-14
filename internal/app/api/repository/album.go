package repository

import (
	"learing_go/simple_crud/internal/app/api/models"
)

type AlbumRepository interface {
	GetAlbums() ([]models.Album, error)
	GetAlbumById(id int64) (*models.Album, error)
	CreateAlbum(album *models.Album) error
	DeleteAlbum(id int64) error
	UpdateAlbum(updatedAlbum models.Album) error
}
