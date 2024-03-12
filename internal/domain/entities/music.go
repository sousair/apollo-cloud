package entities

import (
	"errors"
	"time"
)

type (
	Music struct {
		ID               string    `json:"id"`
		Name             string    `json:"name"`
		DurationInMs     int       `json:"duration_in_ms"`
		ReleaseDate      time.Time `json:"release_date"`
		AlbumID          string    `json:"album_id,omitempty"`
		OwnerID          string    `json:"owner_id"`
		CoverImageURL    string    `json:"cover_image_url,omitempty"`
		SongDataLocation string    `json:"-"`

		Album *Album `json:"album,omitempty"`
		Owner *Owner `json:"owner,omitempty"`
	}

	NewMusicParams struct {
		ID               string
		Name             string
		DurationInMs     int
		ReleaseDate      time.Time
		OwnerID          string
		CoverImageURL    string
		SongDataLocation string

		Album *Album
		Owner *Owner
	}
)

func NewMusic(props NewMusicParams) (*Music, error) {
	// TODO: Put this in a "validate" function
	if props.ID == "" ||
		props.Name == "" ||
		props.DurationInMs == 0 ||
		props.ReleaseDate.IsZero() ||
		props.OwnerID == "" ||
		props.SongDataLocation == "" {
		return nil, errors.New("invalid music")
	}

	music := &Music{
		ID:               props.ID,
		Name:             props.Name,
		DurationInMs:     props.DurationInMs,
		ReleaseDate:      props.ReleaseDate,
		OwnerID:          props.OwnerID,
		CoverImageURL:    props.CoverImageURL,
		SongDataLocation: props.SongDataLocation,
		Album:            props.Album,
		Owner:            props.Owner,
	}

	return music, nil
}
