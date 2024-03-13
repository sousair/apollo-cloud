package entities

import (
	"errors"
	"time"

	"github.com/sousair/apollo-cloud/internal/domain/valueobjects"
)

type (
	Music struct {
		ID           string    `json:"id"`
		Name         string    `json:"name"`
		DurationInMs int       `json:"duration_in_ms"`
		ReleaseDate  time.Time `json:"release_date"`
		AlbumID      string    `json:"album_id,omitempty"`
		OwnerID      string    `json:"owner_id"`

		CoverImageLocation *valueobjects.FileLocation `json:"cover_image_location,omitempty"`
		MusicFileLocation  *valueobjects.FileLocation `json:"-"`

		Album *Album `json:"album,omitempty"`
		Owner *Owner `json:"owner,omitempty"`
	}

	NewMusicParams struct {
		ID           string
		Name         string
		DurationInMs int
		ReleaseDate  time.Time
		AlbumID      string
		OwnerID      string

		CoverImageLocation *valueobjects.FileLocation
		MusicFileLocation  *valueobjects.FileLocation
		Album              *Album
		Owner              *Owner
	}
)

func NewMusic(props NewMusicParams) (*Music, error) {
	// TODO: Put this in a "validate" function
	if props.ID == "" ||
		props.Name == "" ||
		props.DurationInMs == 0 ||
		props.ReleaseDate.IsZero() ||
		props.OwnerID == "" ||
		props.MusicFileLocation == nil {
		return nil, errors.New("invalid music")
	}

	music := &Music{
		ID:                 props.ID,
		Name:               props.Name,
		DurationInMs:       props.DurationInMs,
		ReleaseDate:        props.ReleaseDate,
		AlbumID:            props.AlbumID,
		OwnerID:            props.OwnerID,
		CoverImageLocation: props.CoverImageLocation,
		MusicFileLocation:  props.MusicFileLocation,
		Album:              props.Album,
		Owner:              props.Owner,
	}

	return music, nil
}
