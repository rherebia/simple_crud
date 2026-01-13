package repository

import (
	"learing_go/simple_crud/internal/app/api/models"
)

type AlbumRepository interface {
	GetAlbums() []models.Album
	GetAlbumById(id string) models.Album
	CreateAlbum(album models.Album)
	DeleteAlbum(id string) bool
	UpdateAlbum(updatedAlbum models.Album) bool
}

type AlbumJsonRepository struct {
	albums []models.Album
}

func (r *AlbumJsonRepository) GetAlbums() []models.Album {
	return r.albums
}

func (r *AlbumJsonRepository) GetAlbumById(id string) models.Album {
	for _, album := range r.albums {
		if album.ID == id {
			return album
		}
	}

	return models.Album{}
}

func (r *AlbumJsonRepository) CreateAlbum(album models.Album) {
	r.albums = append(r.albums, album)
}

func (r *AlbumJsonRepository) DeleteAlbum(id string) bool {
	for i, album := range r.albums {
		if album.ID == id {
			r.albums = append(r.albums[:i], r.albums[i+1:]...)
			return true
		}
	}

	return false
}

func (r *AlbumJsonRepository) UpdateAlbum(updatedAlbum models.Album) bool {
	for i, album := range r.albums {
		if album.ID == updatedAlbum.ID {
			r.albums[i] = updatedAlbum
			return true
		}
	}

	return false
}

func NewAlbumJsonRepository() AlbumRepository {
	albums := []models.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return &AlbumJsonRepository{albums: albums}
}
