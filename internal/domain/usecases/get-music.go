package usecases

import "github.com/sousair/apollo-cloud/internal/domain/entities"

type (
	GetMusicUsecase interface {
		Get(params GetMusicParams) (*entities.Music, error)
	}

	GetMusicParams struct {
		ID               string
		IncludeAlbumData bool
		IncludeOwnerData bool
	}
)
