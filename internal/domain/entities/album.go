package entities

import (
	"errors"
	"time"
)

type (
	Album struct {
		ID            string    `json:"id"`
		Name          string    `json:"name"`
		ReleaseDate   time.Time `json:"release_date"`
		OwnerID       string    `json:"owner_id"`
		CoverImageURL string    `json:"cover_image_url"`

		Owner  *Owner   `json:"owner"`
		Musics []*Music `json:"musics"`
	}

	NewAlbumParams struct {
		ID            string
		Name          string
		ReleaseDate   time.Time
		OwnerID       string
		CoverImageURL string
		Owner         *Owner
		Musics        []*Music
	}
)

func NewAlbum(props NewAlbumParams) (*Album, error) {
	if props.ID == "" ||
		props.Name == "" ||
		props.ReleaseDate.IsZero() ||
		props.OwnerID == "" ||
		props.CoverImageURL == "" {
		return nil, errors.New("invalid album")
	}

	album := &Album{
		ID:            props.ID,
		Name:          props.Name,
		ReleaseDate:   props.ReleaseDate,
		OwnerID:       props.OwnerID,
		CoverImageURL: props.CoverImageURL,
		Owner:         props.Owner,
		Musics:        props.Musics,
	}

	return album, nil
}
