package usecases

import "github.com/sousair/apollo-cloud/internal/domain/entities"

type (
	GetAlbumUsecase interface {
		Get(params GetAlbumParams) (*entities.Album, error)
	}

	GetAlbumParams struct {
		ID                string
		IncludeMusicsData bool
		IncludeOwnerData  bool
	}
)
