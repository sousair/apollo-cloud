package entities

import (
	"errors"
	"time"

	"github.com/sousair/apollo-cloud/internal/domain/valueobjects"
)

type (
	Album struct {
		ID                 string                     `json:"id"`
		Name               string                     `json:"name"`
		ReleaseDate        time.Time                  `json:"release_date"`
		OwnerID            string                     `json:"owner_id"`
		CoverImageLocation *valueobjects.FileLocation `json:"cover_image_location,omitempty"`

		Owner  *Owner   `json:"owner,omitempty"`
		Musics []*Music `json:"musics,omitempty"`
	}

	NewAlbumParams struct {
		ID                 string
		Name               string
		ReleaseDate        time.Time
		OwnerID            string
		CoverImageLocation *valueobjects.FileLocation
		Owner              *Owner
		Musics             []*Music
	}
)

func NewAlbum(props NewAlbumParams) (*Album, error) {
	if props.ID == "" ||
		props.Name == "" ||
		props.ReleaseDate.IsZero() ||
		props.OwnerID == "" ||
		props.CoverImageLocation == nil {
		return nil, errors.New("invalid album")
	}

	album := &Album{
		ID:                 props.ID,
		Name:               props.Name,
		ReleaseDate:        props.ReleaseDate,
		OwnerID:            props.OwnerID,
		CoverImageLocation: props.CoverImageLocation,
		Owner:              props.Owner,
		Musics:             props.Musics,
	}

	return album, nil
}
