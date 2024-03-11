package usecases

import (
	"os"

	"github.com/sousair/apollo-cloud/internal/domain/entities"
)

type (
	ReleaseAlbumUsecase interface {
		Release(params ReleaseAlbumParams) (*entities.Album, error)
	}

	ReleaseAlbumParams struct {
		Name           string
		OwnerID        string
		CoverImageFile *os.File
		Musics         []ReleaseAlbumMusicParams
	}

	ReleaseAlbumMusicParams struct {
		Name           string
		DurationInMs   int
		CoverImageFile *os.File
		SongFile       *os.File
	}
)
