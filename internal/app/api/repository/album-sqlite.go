package repository

import (
	"database/sql"
	"learing_go/simple_crud/internal/app/api/models"
)

type AlbumSqliteRepository struct {
	db *sql.DB
}

func (r *AlbumSqliteRepository) GetAlbums() ([]models.Album, error) {
	query := "SELECT * FROM albums"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []models.Album

	for rows.Next() {
		var album models.Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumSqliteRepository) GetAlbumById(id int64) (*models.Album, error) {
	query := "SELECT * FROM albums WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var album models.Album
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	if err != nil {
		return nil, err
	}

	return &album, nil
}

func (r *AlbumSqliteRepository) CreateAlbum(album *models.Album) error {
	query := `
	INSERT INTO albums(title, artist, price)
	VALUES (?, ?, ?)`

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(album.Title, album.Artist, album.Price)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	album.ID = id

	return err
}

func (r *AlbumSqliteRepository) DeleteAlbum(id int64) error {
	query := "DELETE FROM albums WHERE id = ?"

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (r *AlbumSqliteRepository) UpdateAlbum(updatedAlbum models.Album) error {
	query := `
	UPDATE albums
	SET title = ?, artist = ?, price = ?
	WHERE id = ?
	`
	stmt, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, updatedAlbum.ID)
	return err
}

func NewAlbumSqliteRepository(db *sql.DB) AlbumRepository {
	return &AlbumSqliteRepository{db}
}
