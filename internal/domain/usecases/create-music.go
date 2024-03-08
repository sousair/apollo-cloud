package usecases

import (
	"os"
	"time"

	"github.com/sousair/apollo-cloud/internal/domain/entities"
)

type (
	CreateMusicUsecase interface {
		Create(params CreateMusicParams) (*entities.Music, error)
	}

	CreateMusicParams struct {
		Name         string
		OwnerID      string
		AlbumID      string
		DurationInMs int
		ReleaseDate  time.Time
		CoverImage   *os.File
		Song         *os.File
	}
)
