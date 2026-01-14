package repository

import (
	"errors"
	"learing_go/simple_crud/internal/app/api/models"
)

type AlbumJsonRepository struct {
	albums []models.Album
}

func (r *AlbumJsonRepository) GetAlbums() ([]models.Album, error) {
	return r.albums, nil
}

func (r *AlbumJsonRepository) GetAlbumById(id int64) (*models.Album, error) {
	for _, album := range r.albums {
		if album.ID == id {
			return &album, nil
		}
	}

	return nil, errors.New("Album not found")
}

func (r *AlbumJsonRepository) CreateAlbum(newAlbum *models.Album) error {
	var newID int64 = 1

	if len(r.albums) > 0 {
		lastIndex := len(r.albums) - 1
		lastItem := r.albums[lastIndex]
		newID = lastItem.ID + 1
	}

	newAlbum.ID = newID

	r.albums = append(r.albums, *newAlbum)

	return nil
}

func (r *AlbumJsonRepository) DeleteAlbum(id int64) error {
	for i, album := range r.albums {
		if album.ID == id {
			r.albums = append(r.albums[:i], r.albums[i+1:]...)
			return nil
		}
	}

	return errors.New("Album not found")
}

func (r *AlbumJsonRepository) UpdateAlbum(updatedAlbum models.Album) error {
	for i, album := range r.albums {
		if album.ID == updatedAlbum.ID {
			r.albums[i] = updatedAlbum
			return nil
		}
	}

	return errors.New("Album not found")
}

func NewAlbumJsonRepository() AlbumRepository {
	albums := []models.Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return &AlbumJsonRepository{albums: albums}
}
